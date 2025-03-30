package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"
)

var (
	caKey  any
	caCert *x509.Certificate
)

func init() {
	// 解析私钥
	block, _ := pem.Decode(Key)
	if block == nil {
		log.Fatal("解析私钥失败")
	}
	var err error
	caKey, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		// 尝试 PKCS1 格式
		caKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			log.Fatal("解析私钥失败:", err)
		}
	}
	caKey = caKey.(*rsa.PrivateKey)

	// 解析证书
	block, _ = pem.Decode(Cert)
	if block == nil {
		log.Fatal("解析证书失败")
	}
	caCert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal("解析证书失败:", err)
	}
}

var Key = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDBNxo21T9CmgSd
s8bqiS4CY9jnRnwRt7m9/rVJcSYWwj7FYpyZJB0kv9+bDkNbxtbuX+zCjUQMVBUw
cfNa9NA76HPLVhMKq2a9ZPCDjFpjKXG9eqlu62cLALcayHWNUHS5knxFumVulIks
QruB/zh1nxjNHjn/eVegIb2vjgI/J3XKE7zy6dHK602Izzhr6kOt7kHJ2vg73Owz
/iBWYBB3TE38gWpSec3mdcMYRULo92bH7s6Kivn5OQH0DKMbXAN8qjE1wn8mcX+O
K5SrFD/vuQeU/IWJqOWeDXOyCJzCBCDjci4kSAxivOQJzdneG4g1B0uZX/06LW3c
y9PBt3b/AgMBAAECggEAcPr5lyqm7RkGrfibzjUc9ETSZppuniWitRilhSZmsVQw
RFdfXqUhbYXNCOwUMyxYLddlqHXWXqckJMndFUimIwHQBAx24fxL/V9yzFlb2TQ4
0CVTwgpgnKx0ZXLmWhW7y29+v385JKoyc/Epl+BXcV/wrrCIFX0s4fQY6PjPE4HJ
K+w73wlH8Zcp8AtoWJDe7ezBCLKWx8FiDPtnYbQmEAiwDcuXBUtzTcsMdXv+pFEE
buVsi2PHfrgThttWkq0PoZcu/4bflK5eCS2Z5ukC5elO/2TirZRm9msnR5VD9Phf
2LhYBUviJL1Mx2JKceFNpZdA9KXeXFzJAvMQ6UE1cQKBgQDnqi0Ma0u3IaotOpep
wamlzSZl5VOolZlDev5O9iHYBV7Xbe31cqiNudnv/XZhWR6ZzNvmqJE93wsZcD93
NMtF08wNVD1XJkza2vuhj5KeCrtUQAXiq+v8d6bbBzqcTMwQ1KwDkZkGF3SdE3tu
CFD1DQ4aTFZO97iYG+GcsWB6WwKBgQDVgvWrBEoxnzoFhehoQP2QWhiMv7kVFpWu
KG3kcoKMKRg6ubuSlQJLAFRpRdpAAUeRPyV7B45uUHPSbHtNZS9eIL7tGSds1jEM
6iebAUzmX8W6OKBrq8M2mF1knCGZ8QBlxWx/zUUer6Zg0tDn7REid8vUg1sh/IiL
vQNcRUXvLQKBgA1j3K63FO1E2N08Yb/CHdgjf3CsTKWV4AHyffh6aYJe7Rhuli51
riBi+aHReKIv0ID7QkqWu65j0fRqdICjYIxsrYKnt9Pttst/lMPF1Kel3yYDfVOY
tE8na1GFnXGjTrg6UqbyJ3IesPLTSXMWr/c6BfKnLuXpk5XXJrcz2Wu5AoGBAL8L
lGi2yPCHyfIBNMh8Rqawh0EBmEmPUNVp5ZLBB6EEcGKzqGvrmwajP+SGlVgqDPCh
MRdZ1o9Mu6YwXVAOVwkBgfVDSaywJ2mn51JiIn4Mei871gchxOYIBaEttz84jyOB
OKlOcieYAeanHDg6Pte5m5AHDTdm8IMg2G3qqj3tAoGAba6ehRWmExFozo/tVEXx
k2IIii/AIqH/LT0q40fug1aP+tL/2mXwR1ry+IUbVXKgcKMyOJdNDIhLUoAyW1Za
YwDgdciiQeZ+3tl5paT0fuIKQy+3drQcLYs2I6eLbQC8UdkL2WX03v4J/3v/y212
GLnl4EoGVUuJumkN/QVUHdM=
-----END PRIVATE KEY-----`)

var Cert = []byte(`-----BEGIN CERTIFICATE-----
MIIDsjCCApqgAwIBAgIEQC+ajDANBgkqhkiG9w0BAQsFADByMQswCQYDVQQGEwJD
TjEWMBQGA1UEAwwNTUlUTSBQcm94eSBDQTEWMBQGA1UECwwNTUlUTSBQcm94eSBD
QTEVMBMGA1UEBwwMRGVmYXVsdCBDaXR5MRwwGgYDVQQKDBNEZWZhdWx0IENvbXBh
bnkgTHRkMB4XDTI1MDMyNDE0MTMyMFoXDTM1MDMyMjE0MTMyMFowcjELMAkGA1UE
BhMCQ04xFjAUBgNVBAMMDU1JVE0gUHJveHkgQ0ExFjAUBgNVBAsMDU1JVE0gUHJv
eHkgQ0ExFTATBgNVBAcMDERlZmF1bHQgQ2l0eTEcMBoGA1UECgwTRGVmYXVsdCBD
b21wYW55IEx0ZDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAME3GjbV
P0KaBJ2zxuqJLgJj2OdGfBG3ub3+tUlxJhbCPsVinJkkHSS/35sOQ1vG1u5f7MKN
RAxUFTBx81r00Dvoc8tWEwqrZr1k8IOMWmMpcb16qW7rZwsAtxrIdY1QdLmSfEW6
ZW6UiSxCu4H/OHWfGM0eOf95V6Ahva+OAj8ndcoTvPLp0crrTYjPOGvqQ63uQcna
+Dvc7DP+IFZgEHdMTfyBalJ5zeZ1wxhFQuj3ZsfuzoqK+fk5AfQMoxtcA3yqMTXC
fyZxf44rlKsUP++5B5T8hYmo5Z4Nc7IInMIEIONyLiRIDGK85AnN2d4biDUHS5lf
/TotbdzL08G3dv8CAwEAAaNQME4wHQYDVR0OBBYEFDYdU5mibKsy8UvIzTL7CMpT
j39pMB8GA1UdIwQYMBaAFDYdU5mibKsy8UvIzTL7CMpTj39pMAwGA1UdEwQFMAMB
Af8wDQYJKoZIhvcNAQELBQADggEBACGsavN9lUIV3Sblkb5CjQMDIO45k5e8LEkM
VETvldTrq9wA7wb8WPnEdd374R17tmKoPZjZANC8PJPcavDWuC7DNlgMfJN7bfDu
NGy7c8zFL9TaJsUNVXeNmbJ7AjWYlzZES+txRo5PeKUdMmT7Pyhd9L3co65YQofj
hGBKeXkgrQpHiEEvVADl1Zrkf8CtuF1YI3G8Eop4xTv+ix32IXP8VRoAhqyhUzoA
8YCt4yQ6W1bc4gwb0hWu8jzFn6xQ3P1VV3aEWmr8wGi2SFtUk7SEcf+d+rbQp0vY
mcyw5OqOyJhxqDJLr+JosDXJD3DpLa7pPcuZdL7AssPdV2MEoxY=
-----END CERTIFICATE-----`)

// 为指定域名生成证书
func generateCertForHost(host string) (tls.Certificate, error) {
	// 去除端口号
	host = strings.Split(host, ":")[0]

	// 生成私钥
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("生成私钥失败: %v", err)
	}

	// 创建证书模板
	template := &x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject: pkix.Name{
			CommonName: host,
		},
		NotBefore:   time.Now().Add(-time.Hour * 24),
		NotAfter:    time.Now().AddDate(1, 0, 0), // 1年有效期
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:    []string{host},
	}

	// 使用CA证书签发新证书
	certDER, err := x509.CreateCertificate(rand.Reader, template, caCert, &privKey.PublicKey, caKey)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("签发证书失败: %v", err)
	}

	// 将证书转换为PEM格式
	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	})
	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	})

	// 创建tls.Certificate
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("生成X509密钥对失败: %v", err)
	}

	return cert, nil
}

// 添加证书缓存
var (
	certCache     = make(map[string]tls.Certificate)
	certCacheLock sync.RWMutex
)

// 伪造证书加载（使用缓存优化）
func loadFakeCertificate(host string) tls.Certificate {
	// 提取域名部分作为缓存键
	domain := strings.Split(host, ":")[0]

	// 先检查缓存中是否已有该域名的证书
	certCacheLock.RLock()
	cert, exists := certCache[domain]
	certCacheLock.RUnlock()

	if exists {
		log.Printf("使用缓存的证书: %s", domain)
		return cert
	}

	// 缓存中没有，需要生成新证书
	log.Printf("为域名生成新证书: %s", domain)
	cert, err := generateCertForHost(domain)
	if err != nil {
		log.Printf("为域名 %s 生成证书失败: %v", host, err)
		// 使用默认证书作为后备方案
		cert, err = tls.LoadX509KeyPair("server.crt", "server.key")
		if err != nil {
			log.Fatal("默认证书加载失败:", err)
		}
	}

	// 将新生成的证书添加到缓存
	certCacheLock.Lock()
	certCache[domain] = cert
	certCacheLock.Unlock()

	return cert
}
