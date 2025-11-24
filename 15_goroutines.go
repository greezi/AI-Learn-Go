package main

import (
	"fmt"
	"time"
)

/*
15 - Goroutines（协程）

学习目标：
1. 理解 goroutine 的概念
2. 学习如何创建和使用 goroutine
3. 理解并发与并行
4. 学习 goroutine 的同步

运行方式：
go run 15_goroutines.go
*/

func main() {
	fmt.Println("=== Goroutine 基础 ===")

	// 普通函数调用（顺序执行）
	fmt.Println("顺序执行:")
	printNumbers("A")
	printNumbers("B")

	fmt.Println("\n使用 Goroutine（并发执行）:")

	// 使用 go 关键字启动 goroutine
	go printNumbers("1")
	go printNumbers("2")

	// 主 goroutine 需要等待，否则程序会立即退出
	time.Sleep(2 * time.Second)

	fmt.Println("\n=== 匿名函数的 Goroutine ===")

	// 使用匿名函数创建 goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("匿名 goroutine: %d\n", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	// 带参数的匿名 goroutine
	message := "Hello"
	go func(msg string) {
		fmt.Printf("收到消息: %s\n", msg)
	}(message)

	time.Sleep(1 * time.Second)

	fmt.Println("\n=== 多个 Goroutine ===")

	// 启动多个 goroutine
	for i := 1; i <= 5; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d 开始\n", id)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Goroutine %d 结束\n", id)
		}(i)  // 注意：传递 i 作为参数，避免闭包陷阱
	}

	time.Sleep(1 * time.Second)

	fmt.Println("\n=== 闭包陷阱 ===")

	// 错误示例：循环变量闭包
	fmt.Println("错误示例（可能打印相同的数字）:")
	for i := 1; i <= 3; i++ {
		go func() {
			// 所有 goroutine 共享同一个变量 i
			fmt.Printf("错误: %d ", i)
		}()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// 正确示例：传递参数
	fmt.Println("正确示例（传递参数）:")
	for i := 1; i <= 3; i++ {
		go func(n int) {
			// 每个 goroutine 有自己的 n
			fmt.Printf("正确: %d ", n)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	fmt.Println("\n=== Goroutine 与 WaitGroup ===")

	// 使用 sync.WaitGroup 等待所有 goroutine 完成
	// （这里简化演示，详细见后续示例）
	done := make(chan bool)
	count := 3

	for i := 1; i <= count; i++ {
		go func(id int) {
			fmt.Printf("任务 %d 执行中...\n", id)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("任务 %d 完成\n", id)
			done <- true
		}(i)
	}

	// 等待所有 goroutine 完成
	for i := 0; i < count; i++ {
		<-done
	}
	fmt.Println("所有任务完成")

	fmt.Println("\n=== Goroutine 通信示例 ===")

	// goroutine 之间通过 channel 通信
	resultChan := make(chan int)

	go func() {
		sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		resultChan <- sum  // 发送结果
	}()

	result := <-resultChan  // 接收结果
	fmt.Printf("1 到 100 的和: %d\n", result)

	fmt.Println("\n=== 并发计算示例 ===")

	// 并发计算多个任务
	results := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		go func(n int) {
			// 模拟计算
			time.Sleep(100 * time.Millisecond)
			results <- n * n
		}(i)
	}

	// 收集结果
	fmt.Println("平方计算结果:")
	for i := 1; i <= 5; i++ {
		result := <-results
		fmt.Printf("  %d\n", result)
	}

	fmt.Println("\n=== Goroutine 调度 ===")

	// Go 运行时会自动调度 goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("Goroutine A-%d ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("Goroutine B-%d ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	fmt.Println("\n=== 实用示例：并发下载 ===")

	urls := []string{
		"https://example.com/file1",
		"https://example.com/file2",
		"https://example.com/file3",
	}

	finished := make(chan bool, len(urls))

	for i, url := range urls {
		go func(id int, url string) {
			fmt.Printf("开始下载: %s\n", url)
			time.Sleep(300 * time.Millisecond)  // 模拟下载
			fmt.Printf("完成下载: %s\n", url)
			finished <- true
		}(i, url)
	}

	// 等待所有下载完成
	for i := 0; i < len(urls); i++ {
		<-finished
	}
	fmt.Println("所有下载完成")

	fmt.Println("\n=== Goroutine 最佳实践 ===")

	fmt.Println(`
Goroutine 最佳实践：

1. 不要创建过多的 goroutine
   - 每个 goroutine 都有内存开销（约 2KB）
   - 使用工作池模式限制并发数量

2. 总是确保 goroutine 能够退出
   - 避免 goroutine 泄漏
   - 使用 context 管理 goroutine 生命周期

3. 使用 channel 进行通信
   - "不要通过共享内存来通信，而应通过通信来共享内存"

4. 处理 panic
   - goroutine 中的 panic 不会被外部捕获
   - 在 goroutine 内部使用 defer + recover

5. 避免数据竞争
   - 使用 channel 或 sync 包的同步原语
   - 使用 go run -race 检测数据竞争

6. 合理使用缓冲 channel
   - 根据实际需求选择缓冲大小
   - 避免缓冲过大导致内存浪费

7. 注意闭包陷阱
   - 循环中启动 goroutine 时，传递参数而不是使用闭包
	`)

	fmt.Println("\n程序即将退出...")
	time.Sleep(100 * time.Millisecond)
}

// 打印数字的函数
func printNumbers(prefix string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", prefix, i)
		time.Sleep(300 * time.Millisecond)
	}
}
