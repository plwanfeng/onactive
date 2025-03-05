package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Colors 定义ANSI颜色代码
type Colors struct {
	Header    string
	Blue      string
	Green     string
	Yellow    string
	Red       string
	Endc      string
	Bold      string
	Underline string
}

// 初始化颜色常量
var colors = Colors{
	Header:    "\033[95m",
	Blue:      "\033[94m",
	Green:     "\033[92m",
	Yellow:    "\033[93m",
	Red:       "\033[91m",
	Endc:      "\033[0m",
	Bold:      "\033[1m",
	Underline: "\033[4m",
}

// RequestData 请求数据结构
type RequestData struct {
	Code string `json:"code"`
}

func main() {
	// 设置请求头
	headers := map[string]string{
		"Host":                  "orochi.network",
		"Connection":            "keep-alive",
		"sec-ch-ua-platform":    "\"macOS\"",
		"Next-Action":           "0e204ac5ea05e7b92d873b85972e219f2e54a655",
		"sec-ch-ua":             "\"Not(A:Brand\";v=\"99\", \"Google Chrome\";v=\"133\", \"Chromium\";v=\"133\"",
		"sec-ch-ua-mobile":      "?0",
		"Next-Router-State-Tree": "%5B%22%22%2C%7B%22children%22%3A%5B%22(onactive)%22%2C%7B%22children%22%3A%5B%22onactive%22%2C%7B%22children%22%3A%5B%22__PAGE__%22%2C%7B%7D%2C%22%2Fonactive%22%2C%22refresh%22%5D%7D%5D%7D%5D%7D%2Cnull%2Cnull%2Ctrue%5D",
		"User-Agent":            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36",
		"Accept":                "text/x-component",
		"Content-Type":          "text/plain;charset=UTF-8",
		"Origin":                "https://orochi.network",
		"Sec-Fetch-Site":        "same-origin",
		"Sec-Fetch-Mode":        "cors",
		"Sec-Fetch-Dest":        "empty",
		"Referer":               "https://orochi.network/onactive",
		"Accept-Encoding":       "gzip, deflate, br, zstd",
		"Accept-Language":       "zh-CN,zh;q=0.9",
		"Cookie":                "_ga=GA1.1.412046941.1741078396; __Host-onactive-auth.csrf-token=2e10479711b16a8b9e98acbab3ff7fbc5bb372291cf5c7ce39a62f6a1dcb36b2%7Ca82316cffea3289b3e690cb9e76c7f6ae1123d3eb0856054d45814fa2aad3054; __Secure-onactive-auth.callback-url=https%3A%2F%2Forochi.network%2Fonactive; _ga_R9CTTLBRX9=GS1.1.1741137588.2.1.1741138998.0.0.0; __Secure-onactive-auth.session-token=eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIn0..xNklNTOAvu1C8yBy.Dgdnzi2_zdv2zpYEW6leOQbJrBkKsQSAcxH8oo1MhSAHpho1nMoKzh_7R3ncxHH9zTykLMb23QyYoBWjV8yjF_03lMWhYPYAK6TIIHzIS8E5WO7X8gtXhyaonmLxIIez0iqB47nPrGm50V1qcX74espa0CKme2byL4vZrqQUHDJUKtb-0H22jpwQTBVqoFyltNt6aetUyKDC2N9EjlIrprHTMiYcHh4sLHxM8d2QKTG_SSF7umImFKYGqHX1Q_HSgQMK_SXi77Gj6LS7b1nCNfyPJbwuxRZPMxgC-V1K9V329bdzNOsdzHQsb8v3jR9_cTqHkDdT5pDbSZMxZ9PUBQofDXwV0o6xxz7IyLKF0fV_nB3BmcVrr3HhcGLV5QMgS78BCNg1j1pMLCuBxxzNmds9b46jUzxle6MZ86l7fGWls1FMzCL5ZMKKbOb4C5wK5IlsKCrOt8WTUJCtl5EwLWk8kXdmzlL5Ibs6Nme2KKoKOkFk8qs51Mut73PhfKpsL8MH.4AZFrRL5luloJ2SJ_EE9sw",
	}

	// 设置URL
	url := "https://orochi.network/onactive"

	fmt.Printf("%s开始发送POST请求循环，按Ctrl+C退出%s\n", colors.Header, colors.Endc)
	fmt.Printf("%s%s%s\n", colors.Blue, strings.Repeat("=", 50), colors.Endc)

	// 设置信号处理
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 创建reader读取用户输入
	reader := bufio.NewReader(os.Stdin)

	// 主循环
	for {
		select {
		case <-sigChan:
			fmt.Printf("\n%s程序已退出%s\n", colors.Yellow, colors.Endc)
			return
		default:
			// 获取用户输入
			fmt.Print("请输入code参数: ")
			code, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("%s读取输入错误: %v%s\n", colors.Red, err, colors.Endc)
				continue
			}

			// 去除输入中的空白字符
			code = strings.TrimSpace(code)

			// 准备请求数据
			data := []RequestData{{Code: code}}
			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Printf("%s生成JSON数据错误: %v%s\n", colors.Red, err, colors.Endc)
				continue
			}

			// 创建请求
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Printf("%s创建请求错误: %v%s\n", colors.Red, err, colors.Endc)
				continue
			}

			// 设置请求头
			for key, value := range headers {
				req.Header.Set(key, value)
			}

			// 发送请求
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("%s请求发生错误: %v%s\n", colors.Red, err, colors.Endc)
				fmt.Printf("%s%s%s\n", colors.Blue, strings.Repeat("=", 50), colors.Endc)
				continue
			}
			defer resp.Body.Close()

			// 打印时间戳
			fmt.Printf("\n%s请求时间: %s%s\n", colors.Yellow, time.Now().Format("2006-01-02 15:04:05"), colors.Endc)

			// 读取响应内容
			var reader io.ReadCloser
			switch resp.Header.Get("Content-Encoding") {
			case "gzip":
				reader, err = gzip.NewReader(resp.Body)
				if err != nil {
					fmt.Printf("%s解压gzip数据错误: %v%s\n", colors.Red, err, colors.Endc)
					continue
				}
				defer reader.Close()
			default:
				reader = resp.Body
			}

			body, err := io.ReadAll(reader)
			if err != nil {
				fmt.Printf("%s读取响应错误: %v%s\n", colors.Red, err, colors.Endc)
				continue
			}

			// 打印响应内容
			fmt.Printf("%s响应内容:%s\n", colors.Blue, colors.Endc)
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, body, "", "  "); err != nil {
				// 尝试将响应作为UTF-8字符串处理
				fmt.Printf("%s%s%s\n", colors.Red, string(body), colors.Endc)
			} else {
				fmt.Println(prettyJSON.String())
			}

			fmt.Printf("%s%s%s\n", colors.Blue, strings.Repeat("=", 50), colors.Endc)
		}
	}
}