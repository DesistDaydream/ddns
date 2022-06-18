package main

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	// 发起 HTTP GET 请求
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		logrus.Fatalln(err)
	}

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	logrus.Infoln(string(body), err)
}
