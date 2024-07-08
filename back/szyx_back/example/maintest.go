package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// 模拟生成大模型数据的函数
func generateData(ch chan<- string) {
	defer close(ch)
	lines := []string{
		"Hello, this is the first line.\n",
		"Now we are on the second line.\n",
		"Finally, here is the third line.\n",
	}

	for _, line := range lines {
		ch <- line
		time.Sleep(1 * time.Second) // 模拟延迟
	}
}

// 处理HTTP请求
func streamHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string) // 创建Channel来传输数据

	// 启动一个Goroutine来生成数据
	//go generateData(ch)
	content, _ := io.ReadAll(r.Body)
	fmt.Println(content)
	//go kdxf.Test(string(content), ch)
	// 设置响应头为text/plain和chunked
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Transfer-Encoding", "chunked")

	// 使用Flusher强制向客户端刷新数据
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// 逐行从Channel读取数据并写入响应
	for line := range ch {
		w.Write([]byte(line))
		//fmt.Fprintf(w, line)
		flusher.Flush()
	}
}

// 处理HTTP请求
func batchHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string) // 创建Channel来传输数据

	// 启动一个Goroutine来生成数据
	//go generateData(ch)
	content, _ := io.ReadAll(r.Body)
	fmt.Println(content)
	//go kdxf.Test(string(content), ch)
	// 设置响应头为text/plain和chunked
	w.Header().Set("Content-Type", "text/plain")

	var result string
	// 逐行从Channel读取数据并写入响应
	for line := range ch {
		result = result + line
	}
	fmt.Println(result)
	w.Write([]byte(result))
}

// 处理HTML页面请求
func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/test.html")
}
func main() {
	http.HandleFunc("/stream", streamHandler)
	http.HandleFunc("/batch", batchHandler)
	http.HandleFunc("/", htmlHandler) // 添加HTML处理函数
	fmt.Println("Starting server at :9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
