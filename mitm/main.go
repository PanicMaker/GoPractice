package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// 1. 启动代理监听
func main() {
	listener, err := net.Listen("tcp", ":8443")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MITM代理运行在 :8443")

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go handleClient(clientConn)
	}
}

// 2. 处理客户端连接
func handleClient(clientConn net.Conn) {
	defer clientConn.Close()

	// 读取客户端请求（假设是HTTPS CONNECT请求）
	buf := make([]byte, 1024)
	n, err := clientConn.Read(buf)
	if err != nil {
		log.Println("Read error:", err)
		return
	}

	// 解析CONNECT请求中的目标域名
	request := string(buf[:n])
	targetHost, err := extractHostFromCONNECT(request)
	if err != nil {
		log.Println("解析CONNECT请求失败:", err)
		return
	}

	// 向客户端发送200 Connection Established
	_, err = clientConn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
	if err != nil {
		log.Println("发送Connection Established失败:", err)
		return
	}

	// 3. 连接到真实目标服务器
	log.Printf("正在连接目标服务器: %s", targetHost)
	targetConn, err := tls.Dial("tcp", targetHost, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Println("连接目标服务器失败:", err)
		return
	}
	defer targetConn.Close()

	// 4. 向客户端发送伪造的TLS证书（关键步骤）
	// 此处需动态生成目标域名的证书（示例中省略证书生成逻辑）
	clientTLSConfig := &tls.Config{
		Certificates: []tls.Certificate{loadFakeCertificate(targetHost)}, // 加载伪造证书
	}

	// 5. 与客户端建立TLS连接（劫持）
	clientTLS := tls.Server(clientConn, clientTLSConfig)
	err = clientTLS.Handshake()
	if err != nil {
		log.Printf("TLS握手失败: %v (目标: %s)", err, targetHost)
		// 添加更多调试信息
		log.Printf("TLS配置: 证书数量=%d", len(clientTLSConfig.Certificates))
		return
	}
	defer clientTLS.Close()

	// 6. 双向转发解密后的数据
	go io.Copy(targetConn, clientTLS)
	io.Copy(clientTLS, targetConn)
}



// 从CONNECT请求中提取目标域名
func extractHostFromCONNECT(request string) (string, error) {
	// CONNECT请求格式: CONNECT example.com:443 HTTP/1.1
	parts := strings.Split(request, " ")
	if len(parts) < 2 || parts[0] != "CONNECT" {
		return "", fmt.Errorf("无效的CONNECT请求")
	}

	// 返回域名部分 (example.com:443)
	return parts[1], nil
}
