package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 测试数据模型
type TestItem struct {
	ID    uint    `gorm:"primarykey" json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	gorm.Model
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database connection")
	}

	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置连接的最大生命周期
	sqlDB.SetConnMaxLifetime(time.Hour)

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
				Name:  "测试数据-" + strconv.Itoa(i+j),
				Value: float64(i+j) * 1.5,
				Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			})
		}
		db.CreateInBatches(items, batchSize)

		if (i+batchSize)%100000 == 0 {
			println("已插入", i+batchSize, "条数据")
		}
	}
}

func main() {
	db := initDB()

	// 如果需要插入测试数据，取消下面这行的注释
	// insertTestData(db)

	r := gin.Default()

	r.GET("/api/items", func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
		lastID, _ := strconv.Atoi(c.DefaultQuery("last_id", "0"))

		var total int64
		var items []TestItem

		// 获取总记录数（可以考虑缓存这个值）
		db.Model(&TestItem{}).Count(&total)

		// 使用 ID 作为游标进行查询
		query := db.Model(&TestItem{})
		if lastID > 0 {
			query = query.Where("id > ?", lastID)
		}
		query.Order("id asc").Limit(pageSize).Find(&items)

		c.JSON(http.StatusOK, PageResponse{
			Data:     items,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		})
	})

	r.Run(":8080")
}
