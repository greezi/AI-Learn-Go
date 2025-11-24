package main

import (
	"fmt"
	"time"
)

/*
16 - Channels（通道）

学习目标：
1. 理解 channel 的概念和用途
2. 学习 channel 的创建和使用
3. 学习有缓冲和无缓冲 channel
4. 学习 channel 的关闭和遍历
5. 学习 select 语句

运行方式：
go run 16_channels.go
*/

func main() {
	fmt.Println("=== Channel 基础 ===")

	// 创建 channel
	ch := make(chan int)

	// 在 goroutine 中发送数据
	go func() {
		fmt.Println("发送: 42")
		ch <- 42  // 发送数据到 channel
	}()

	// 从 channel 接收数据
	value := <-ch
	fmt.Printf("接收: %d\n\n", value)

	fmt.Println("=== 无缓冲 Channel ===")

	// 无缓冲 channel：发送操作会阻塞，直到有接收者
	unbuffered := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("准备接收...")
		msg := <-unbuffered
		fmt.Printf("接收到: %s\n", msg)
	}()

	fmt.Println("准备发送...")
	unbuffered <- "Hello"  // 会阻塞直到有接收者
	fmt.Println("发送完成\n")

	fmt.Println("=== 有缓冲 Channel ===")

	// 有缓冲 channel：可以存储一定数量的值
	buffered := make(chan int, 3)  // 缓冲区大小为 3

	// 发送不会阻塞，直到缓冲区满
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println("发送了 3 个值到缓冲 channel")

	// 接收值
	fmt.Printf("接收: %d\n", <-buffered)
	fmt.Printf("接收: %d\n", <-buffered)
	fmt.Printf("接收: %d\n\n", <-buffered)

	fmt.Println("=== Channel 方向 ===")

	// 只发送 channel
	sendOnly := make(chan int)
	go sender(sendOnly)
	fmt.Printf("从只发送 channel 接收: %d\n", <-sendOnly)

	// 只接收 channel
	receiveOnly := make(chan int)
	go func() {
		receiveOnly <- 100
	}()
	receiver(receiveOnly)
	fmt.Println()

	fmt.Println("=== 关闭 Channel ===")

	ch2 := make(chan int, 3)
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	close(ch2)  // 关闭 channel

	// 从已关闭的 channel 接收数据仍然可以
	fmt.Printf("接收: %d\n", <-ch2)
	fmt.Printf("接收: %d\n", <-ch2)
	fmt.Printf("接收: %d\n", <-ch2)

	// 从已关闭且为空的 channel 接收会得到零值
	v, ok := <-ch2
	fmt.Printf("接收: %d, channel 是否打开: %v\n\n", v, ok)

	fmt.Println("=== Range 遍历 Channel ===")

	ch3 := make(chan int, 5)

	// 发送数据
	go func() {
		for i := 1; i <= 5; i++ {
			ch3 <- i
		}
		close(ch3)  // 必须关闭，否则 range 会一直等待
	}()

	// 使用 range 遍历 channel
	fmt.Println("遍历 channel:")
	for value := range ch3 {
		fmt.Printf("  接收: %d\n", value)
	}
	fmt.Println()

	fmt.Println("=== Select 语句 ===")

	// select 用于处理多个 channel 操作
	ch4 := make(chan string)
	ch5 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch4 <- "来自 ch4"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch5 <- "来自 ch5"
	}()

	// select 会选择第一个准备好的 channel
	select {
	case msg1 := <-ch4:
		fmt.Println(msg1)
	case msg2 := <-ch5:
		fmt.Println(msg2)
	}
	fmt.Println()

	fmt.Println("=== Select 与 Default ===")

	ch6 := make(chan int)

	select {
	case val := <-ch6:
		fmt.Printf("接收到: %d\n", val)
	default:
		fmt.Println("没有数据可接收，执行默认操作")
	}
	fmt.Println()

	fmt.Println("=== Select 超时处理 ===")

	ch7 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch7 <- "延迟消息"
	}()

	select {
	case msg := <-ch7:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("超时：1 秒内没有收到消息")
	}
	fmt.Println()

	fmt.Println("=== 工作池模式 ===")

	// 创建任务和结果 channel
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// 启动 3 个工作者
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("结果: %d\n", result)
	}
	fmt.Println()

	fmt.Println("=== Channel 同步 ===")

	done := make(chan bool)

	go func() {
		fmt.Println("执行任务...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("任务完成")
		done <- true
	}()

	fmt.Println("等待任务完成...")
	<-done
	fmt.Println("主程序继续执行\n")

	fmt.Println("=== 生产者-消费者模式 ===")

	// 数据 channel
	data := make(chan int, 5)

	// 生产者
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("生产: %d\n", i)
			data <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()

	// 消费者
	go func() {
		for value := range data {
			fmt.Printf("  消费: %d\n", value)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println()

	fmt.Println("=== Fan-Out / Fan-In 模式 ===")

	// Fan-Out：一个输入，多个处理者
	input := make(chan int, 10)
	output := make(chan int, 10)

	// 发送输入
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()

	// 启动多个处理者（Fan-Out）
	for i := 1; i <= 3; i++ {
		go func(id int) {
			for num := range input {
				fmt.Printf("处理者 %d 处理: %d\n", id, num)
				output <- num * 2
			}
		}(i)
	}

	// 等待处理完成
	go func() {
		time.Sleep(1 * time.Second)
		close(output)
	}()

	// 收集所有结果（Fan-In）
	fmt.Println("处理结果:")
	for result := range output {
		fmt.Printf("  %d\n", result)
	}
	fmt.Println()

	fmt.Println("=== Channel 最佳实践 ===")

	fmt.Println(`
Channel 最佳实践：

1. 谁创建谁关闭
   - 发送者负责关闭 channel
   - 接收者不应该关闭 channel

2. 关闭 channel 的注意事项
   - 向已关闭的 channel 发送数据会 panic
   - 关闭已关闭的 channel 会 panic
   - 从已关闭的 channel 接收数据安全

3. 使用 range 遍历 channel
   - 自动处理 channel 关闭
   - 代码更简洁

4. 合理使用缓冲
   - 无缓冲：需要发送和接收同步
   - 有缓冲：减少阻塞，提高性能
   - 根据实际需求选择缓冲大小

5. 使用 select 处理多个 channel
   - 超时控制
   - 非阻塞操作（default）
   - 多路复用

6. 避免 channel 泄漏
   - 确保所有发送的数据都被接收
   - 使用 context 控制 goroutine 生命周期

7. nil channel 的行为
   - 向 nil channel 发送数据会永久阻塞
   - 从 nil channel 接收数据会永久阻塞
   - 在 select 中可以利用这个特性
	`)
}

// 只发送 channel 参数
func sender(ch chan<- int) {
	ch <- 42
}

// 只接收 channel 参数
func receiver(ch <-chan int) {
	value := <-ch
	fmt.Printf("接收者收到: %d\n", value)
}

// 工作者函数
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("工作者 %d 处理任务 %d\n", id, job)
		time.Sleep(300 * time.Millisecond)
		results <- job * 2
	}
}
