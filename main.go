package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type result struct {
	uid int
	m3b string
}

type tokenRes struct {
	Token   string `json:"token"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	file, err := os.Open("./account.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	results := []result{}
	reg := regexp.MustCompile(`"m3b":(.+?),`)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		s := strings.Split(str, ",")
		uid, err := strconv.Atoi(strings.TrimSpace(s[len(s)-1]))
		if err != nil {
			panic(err)
		}
		fmt.Println(uid)

		token := getToken(uid)
		// fmt.Println(token)

		body := request(uid, token)
		// fmt.Println(body)
		m3b := reg.FindStringSubmatch(body)
		// fmt.Println(m3b)
		tmp := result{}
		if m3b != nil {
			tmp = result{
				uid: uid,
				m3b: m3b[1],
			}
		} else {
			tmp = result{
				uid: uid,
				m3b: "",
			}
		}

		results = append(results, tmp)
	}
	save(results)
	// fmt.Println(results)
}

func getToken(uid int) string {
	var dataUrl string = "http://action.op.dianhun.cn/202112/niubi/public/index.php/api/user/loginTest?uid=" + strconv.Itoa(uid)
	resp, err := http.Get(dataUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	var res tokenRes
	_ = json.Unmarshal(body, &res)
	// fmt.Printf("%+v", res)
	return res.Token
}

func request(uid int, token string) string {
	client := &http.Client{}
	url := "http://action.op.dianhun.cn/202112/niubi/public/index.php/api/user/bind"
	req, err := http.NewRequest("POST", url, strings.NewReader("m3AreaId=110"))
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func save(results []result) {
	csvFile, err := os.Create("m3b.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	// 初始化一个 csv writer，并通过这个 writer 写入数据到 csv 文件
	writer := csv.NewWriter(csvFile)
	for _, v := range results {
		line := []string{
			strconv.Itoa(v.uid), // 将 int 类型数据转化为字符串
			v.m3b,
		}
		// 将切片类型行数据写入 csv 文件
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	// 将 writer 缓冲中的数据都推送到 csv 文件，至此就完成了数据写入到 csv 文件
	writer.Flush()
}
