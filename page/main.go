package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 测试数据模型
type TestItem struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `json:"name"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

// 分页响应结构
type PageResponse struct {
	Data     []TestItem `json:"data"`
	Total    int64      `json:"total"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
}

// 初始化数据库连接
func initDB() *gorm.DB {
	dsn := "root:1234567890@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 添加性能优化配置
		PrepareStmt: true,                                  // 缓存预编译语句
		Logger:      logger.Default.LogMode(logger.Silent), // 关闭SQL日志，提高性能
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database connection")
	}

	// 增加连接池配置以支持高并发
	sqlDB.SetMaxIdleConns(50)  // 增加空闲连接数
	sqlDB.SetMaxOpenConns(200) // 增加最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 30) // 设置空闲连接超时

	// 自动迁移表结构
	db.AutoMigrate(&TestItem{})
	return db
}

// 插入测试数据
func insertTestData(db *gorm.DB) {
	batchSize := 1000
	totalRecords := 10000000

	for i := 0; i < totalRecords; i += batchSize {
		var items []TestItem
		for j := 0; j < batchSize && (i+j) < totalRecords; j++ {
			items = append(items, TestItem{
				Name:      "测试数据-" + strconv.Itoa(i+j),
				Value:     float64(i+j) * 1.5,
				CreatedAt: time.Now(),
			})
		}
		db.CreateInBatches(items, batchSize)

		if (i+batchSize)%100000 == 0 {
			println("已插入", i+batchSize, "条数据")
		}
	}
}

// 添加一个全局变量存储总记录数和上次更新时间
var (
	cachedTotal   int64
	lastCountTime time.Time
	cacheMutex    sync.RWMutex
)

// 异步更新总记录数的函数
func updateTotalCountAsync(db *gorm.DB) {
	for {
		var total int64
		db.Model(&TestItem{}).Count(&total)

		cacheMutex.Lock()
		cachedTotal = total
		lastCountTime = time.Now()
		cacheMutex.Unlock()

		time.Sleep(time.Minute * 5) // 每5分钟更新一次
	}
}

func main() {
	db := initDB()

	// 如果需要插入测试数据，取消下面这行的注释
	// insertTestData(db)

	// 启动异步更新总记录数的协程
	go updateTotalCountAsync(db)

	r := gin.Default()

	// 添加性能优化中间件
	r.Use(gin.Recovery())

	// 使用更高效的游标分页API
	r.GET("/api/items/cursor", func(c *gin.Context) {
		pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
		lastID, _ := strconv.Atoi(c.DefaultQuery("last_id", "0"))

		var items []TestItem

		// 使用游标分页，避免OFFSET
		query := db.Model(&TestItem{})
		if lastID > 0 {
			query = query.Where("id > ?", lastID)
		}
		query.Order("id asc").Limit(pageSize + 1).Find(&items)

		hasMore := false
		if len(items) > pageSize {
			hasMore = true
			items = items[:pageSize] // 移除多查询的那一条
		}

		nextLastID := uint(0)
		if len(items) > 0 {
			nextLastID = items[len(items)-1].ID
		}

		cacheMutex.RLock()
		total := cachedTotal
		cacheMutex.RUnlock()

		c.JSON(http.StatusOK, gin.H{
			"data":         items,
			"has_more":     hasMore,
			"next_last_id": nextLastID,
			"total":        total,
		})
	})

	// 保留原有API但优化性能
	r.GET("/api/items", func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

		var items []TestItem

		// 使用优化的查询方式
		if page > 1 {
			// 对于非第一页，使用ID范围查询代替OFFSET
			var firstID uint
			db.Model(&TestItem{}).Order("id asc").
				Offset((page-1)*pageSize).Limit(1).
				Pluck("id", &firstID)

			if firstID > 0 {
				db.Where("id >= ?", firstID).
					Order("id asc").Limit(pageSize).
					Find(&items)
			}
		} else {
			// 第一页直接查询
			db.Order("id asc").Limit(pageSize).Find(&items)
		}

		// 使用缓存的总记录数
		cacheMutex.RLock()
		total := cachedTotal
		cacheMutex.RUnlock()

		c.JSON(http.StatusOK, PageResponse{
			Data:     items,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		})
	})

	r.Run(":8080")
}
