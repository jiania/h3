package main

import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "github.com/quic-go/quic-go/http3"
)

func main() {
    // 检查命令行参数
    if len(os.Args) != 2 {
        fmt.Println("Usage: quic [url]")
        os.Exit(1)
    }
    url := os.Args[1] // 获取 URL

    // 创建支持 HTTP/3 的客户端
    client := &http.Client{
        Transport: &http3.RoundTripper{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true, // 仅用于测试，避免在生产环境中使用
            },
        },
    }
    defer client.CloseIdleConnections()

    // 发送 HTTP/3 请求
    resp, err := client.Get(url)
    if err != nil {
        fmt.Println("Error making request:", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    // 读取响应内容
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        os.Exit(1)
    }

    fmt.Printf("%s", body)
}
