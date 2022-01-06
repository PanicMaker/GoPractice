package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 接口请求URL
	apiUrl := "http://localhost" // 不要使用接口地址测试
	max := 1<<63 - 1             // 最大整数，模拟大量数据

	// 初始化参数
	param := url.Values{}
	// 配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("a", "test") // 需要查询的手机号码或手机号码前7位

	for i := 0; i < max; i++ {
		go func(i int) {
			// 一些逻辑代码...
			fmt.Printf("start func: %dn", i)
			// 发送请求
			data, err := Get(apiUrl, param)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(data))
		}(i)
	}
}

// Get 方式发起网络请求
func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:rn%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
