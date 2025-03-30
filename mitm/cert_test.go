package main

import (
	"sync"
	"testing"
	"time"
)

// 测试证书生成性能
func BenchmarkLoadFakeCertificate(b *testing.B) {
	// 重置计时器
	b.ResetTimer()

	// 运行 b.N 次证书生成
	for i := 0; i < b.N; i++ {
		loadFakeCertificate("example.com:443")
	}
}

// 测试相同域名重复生成证书的性能
func BenchmarkLoadFakeCertificateRepeated(b *testing.B) {
	domain := "example.com:443"

	// 重置计时器
	b.ResetTimer()

	// 重复为同一域名生成证书
	for i := 0; i < b.N; i++ {
		loadFakeCertificate(domain)
	}
}

// 测试不同域名生成证书的性能
func BenchmarkLoadFakeCertificateDifferentDomains(b *testing.B) {
	domains := []string{
		"example.com:443",
		"google.com:443",
		"github.com:443",
		"baidu.com:443",
		"microsoft.com:443",
	}

	// 重置计时器
	b.ResetTimer()

	// 为不同域名生成证书
	for i := 0; i < b.N; i++ {
		loadFakeCertificate(domains[i%len(domains)])
	}
}

// 测试并发生成证书的性能
func BenchmarkLoadFakeCertificateConcurrent(b *testing.B) {
	domains := []string{
		"example.com:443",
		"google.com:443",
		"github.com:443",
		"baidu.com:443",
		"microsoft.com:443",
	}

	// 创建等待组
	var wg sync.WaitGroup

	// 重置计时器
	b.ResetTimer()

	// 并发生成证书
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			loadFakeCertificate(domains[idx%len(domains)])
		}(i)
	}

	// 等待所有goroutine完成
	wg.Wait()
}

// 测试证书生成的实际耗时
func TestCertificateGenerationTime(t *testing.T) {
	domains := []string{
		"example.com:443",
		"google.com:443",
		"github.com:443",
		"baidu.com:443",
		"microsoft.com:443",
	}

	for _, domain := range domains {
		start := time.Now()
		loadFakeCertificate(domain)
		elapsed := time.Since(start)
		t.Logf("为域名 %s 生成证书耗时: %v", domain, elapsed)
	}
}
