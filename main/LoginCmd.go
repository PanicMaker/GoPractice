package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	var host string
	var user string
	var password string

	flag.StringVar(&user, "u", "admin", "账户，默认为admin")
	flag.StringVar(&password, "p", "1Syscore#", "密码，默认为1Sysc0r#")
	flag.StringVar(&host, "h", "http://isc-permission-service:32100/api/permission/auth/login", "用户登录地址:默认为http://isc-permission-service:32100/api/permission/auth/login")
	//命令解析
	flag.Parse()
	req := "{\"loginName\":\"" + user +
		"\",\"password\":\"" + password +
		"\"}"
	resp, err := http.Post(host, "application/json", strings.NewReader(req))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}
