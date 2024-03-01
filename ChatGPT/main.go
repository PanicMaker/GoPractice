package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {

	openAiUrl := "https://api.openai.com/v1/chat/completions"
	method := "POST"

	payload := strings.NewReader(`{
  "model": "gpt-3.5-turbo",
  "messages": [{"role": "user", "content": "Hello!"}]
}`)

	proxyURL, err := url.Parse("http://127.0.0.1:7890")
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(method, openAiUrl, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1UaEVOVUpHTkVNMVFURTRNMEZCTWpkQ05UZzVNRFUxUlRVd1FVSkRNRU13UmtGRVFrRXpSZyJ9.eyJodHRwczovL2FwaS5vcGVuYWkuY29tL3Byb2ZpbGUiOnsiZW1haWwiOiJrYXRoaWV6aGJyYUBob3RtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlfSwiaHR0cHM6Ly9hcGkub3BlbmFpLmNvbS9hdXRoIjp7InBvaWQiOiJvcmctVFY1Qk9IRkRiN1hHT09kSWc5Q3BKME90IiwidXNlcl9pZCI6InVzZXItaFdsZ3V0Wmc0dHp0TGZscDFlbnAyZmlqIn0sImlzcyI6Imh0dHBzOi8vYXV0aDAub3BlbmFpLmNvbS8iLCJzdWIiOiJhdXRoMHw2M2Q4ZWIzN2RmMmE4ZTZiMTZhYTM2NGYiLCJhdWQiOlsiaHR0cHM6Ly9hcGkub3BlbmFpLmNvbS92MSIsImh0dHBzOi8vb3BlbmFpLm9wZW5haS5hdXRoMGFwcC5jb20vdXNlcmluZm8iXSwiaWF0IjoxNzA1OTEzMTcyLCJleHAiOjE3MDY3NzcxNzIsImF6cCI6IlRkSkljYmUxNldvVEh0Tjk1bnl5d2g1RTR5T282SXRHIiwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSBtb2RlbC5yZWFkIG1vZGVsLnJlcXVlc3Qgb3JnYW5pemF0aW9uLnJlYWQgb3JnYW5pemF0aW9uLndyaXRlIG9mZmxpbmVfYWNjZXNzIn0.TtPJ03_lZHtjDxRStaUGKhVLQrgZbDg0gnkjtMjSFpa-V2WD-r2ZeoicEgSiTy3wjCBgAD8gLtb6aa5MP7PGxEoRr0pIcdW1nU-80HpcT-L5Uo0DpEzbtLMcSKT_AsikDuYYt5vJzMWg2cEpQvwVcGvaS917Y7E7KScIZyfEVdhmIgN_07hMZFCYbp6mm_nzMG1SkZVxByvTTsTYRL2dphKToLSwa-OL2hvkrB1o511p1wwMeBUrg15mtSnaTJe0awHDi9uwWMgo8WFMSX5Cgbnp1PRbbTXh7apx-gAmT2VsFUGvru-LcvYny2SA1XELl1kRD-wheP79tATuP4krew")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
