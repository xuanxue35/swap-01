package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Swap1 服务已启动，监听端口：8080")
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "欢迎访问 Swap1 服务")
}
