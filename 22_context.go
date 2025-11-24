package main

import (
	"context"
	"fmt"
	"time"
)

/*
22 - Context

学习目标：
1. 理解 Context 的用途
2. 学习 Context 的创建方式
3. 学习 Context 的传播
4. 掌握 Context 的使用场景

运行方式：
go run 22_context.go
*/

func main() {
	fmt.Println("=== Context 基础 ===")

	// context.Background() - 根 context
	ctx := context.Background()
	fmt.Printf("Background context: %v\n", ctx)

	// context.TODO() - 不确定使用哪种 context 时的占位符
	todoCtx := context.TODO()
	fmt.Printf("TODO context: %v\n", todoCtx)

	fmt.Println("\n=== context.WithCancel ===")

	// 创建可取消的 context
	ctx1, cancel := context.WithCancel(context.Background())

	// 启动 goroutine
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("  goroutine 收到取消信号，退出")
				return
			default:
				fmt.Println("  goroutine 工作中...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(ctx1)

	// 等待一段时间后取消
	time.Sleep(500 * time.Millisecond)
	cancel()
	fmt.Println("已发送取消信号")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n=== context.WithTimeout ===")

	// 创建带超时的 context
	ctx2, cancel2 := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel2()  // 总是调用 cancel 释放资源

	// 模拟慢操作
	result := make(chan string, 1)
	go func() {
		time.Sleep(300 * time.Millisecond)  // 模拟操作
		result <- "操作完成"
	}()

	select {
	case r := <-result:
		fmt.Printf("结果: %s\n", r)
	case <-ctx2.Done():
		fmt.Printf("超时: %v\n", ctx2.Err())
	}

	// 超时的情况
	ctx3, cancel3 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel3()

	result2 := make(chan string, 1)
	go func() {
		time.Sleep(200 * time.Millisecond)  // 操作时间超过超时时间
		result2 <- "操作完成"
	}()

	select {
	case r := <-result2:
		fmt.Printf("结果: %s\n", r)
	case <-ctx3.Done():
		fmt.Printf("超时: %v\n", ctx3.Err())
	}

	fmt.Println("\n=== context.WithDeadline ===")

	// 创建带截止时间的 context
	deadline := time.Now().Add(500 * time.Millisecond)
	ctx4, cancel4 := context.WithDeadline(context.Background(), deadline)
	defer cancel4()

	fmt.Printf("截止时间: %v\n", deadline)

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("操作完成")
	case <-ctx4.Done():
		fmt.Printf("已到达截止时间: %v\n", ctx4.Err())
	}

	fmt.Println("\n=== context.WithValue ===")

	// 创建带值的 context
	type contextKey string
	const userIDKey contextKey = "userID"
	const requestIDKey contextKey = "requestID"

	ctx5 := context.WithValue(context.Background(), userIDKey, 12345)
	ctx5 = context.WithValue(ctx5, requestIDKey, "req-abc-123")

	// 获取值
	if userID, ok := ctx5.Value(userIDKey).(int); ok {
		fmt.Printf("用户 ID: %d\n", userID)
	}
	if reqID, ok := ctx5.Value(requestIDKey).(string); ok {
		fmt.Printf("请求 ID: %s\n", reqID)
	}

	// 在函数链中传递 context
	processRequest(ctx5)

	fmt.Println("\n=== Context 传播 ===")

	// Context 应该在函数间传递
	ctx6, cancel6 := context.WithCancel(context.Background())

	fmt.Println("启动任务链...")
	go task1(ctx6)

	time.Sleep(300 * time.Millisecond)
	cancel6()
	fmt.Println("取消所有任务")
	time.Sleep(200 * time.Millisecond)

	fmt.Println("\n=== 实用示例：HTTP 请求超时 ===")

	simulateHTTPRequest()

	fmt.Println("\n=== Context 最佳实践 ===")

	fmt.Println(`
Context 最佳实践：

1. 将 Context 作为函数第一个参数
   func DoSomething(ctx context.Context, arg Arg) error

2. 不要将 Context 存储在结构体中
   ✗ type Service struct { ctx context.Context }
   ✓ func (s *Service) Do(ctx context.Context) error

3. 传递 Context，不要传递 nil
   ✗ DoSomething(nil, arg)
   ✓ DoSomething(context.Background(), arg)

4. 使用 context.Value 只传递请求范围的值
   - 请求 ID
   - 认证令牌
   - 跟踪 ID
   不要用于传递可选参数

5. 总是调用 cancel 函数
   ctx, cancel := context.WithTimeout(...)
   defer cancel()

6. 使用自定义类型作为 key
   type contextKey string
   const myKey contextKey = "myKey"

7. 检查 ctx.Done() 进行取消处理
   select {
   case <-ctx.Done():
       return ctx.Err()
   default:
       // 继续处理
   }

8. Context 是只读的
   - 不要修改传入的 Context
   - 需要新值时创建派生 Context
	`)
}

// 处理请求示例
func processRequest(ctx context.Context) {
	type contextKey string
	const userIDKey contextKey = "userID"
	const requestIDKey contextKey = "requestID"

	fmt.Println("\n处理请求:")
	if userID := ctx.Value(userIDKey); userID != nil {
		fmt.Printf("  当前用户: %v\n", userID)
	}
	if reqID := ctx.Value(requestIDKey); reqID != nil {
		fmt.Printf("  请求 ID: %v\n", reqID)
	}
}

// 任务链示例
func task1(ctx context.Context) {
	fmt.Println("  任务 1 开始")

	select {
	case <-ctx.Done():
		fmt.Println("  任务 1 被取消")
		return
	case <-time.After(100 * time.Millisecond):
		task2(ctx)  // 传递 context
	}
}

func task2(ctx context.Context) {
	fmt.Println("  任务 2 开始")

	select {
	case <-ctx.Done():
		fmt.Println("  任务 2 被取消")
		return
	case <-time.After(100 * time.Millisecond):
		task3(ctx)
	}
}

func task3(ctx context.Context) {
	fmt.Println("  任务 3 开始")

	select {
	case <-ctx.Done():
		fmt.Println("  任务 3 被取消")
		return
	case <-time.After(500 * time.Millisecond):
		fmt.Println("  任务 3 完成")
	}
}

// 模拟 HTTP 请求
func simulateHTTPRequest() {
	fmt.Println("模拟 HTTP 请求处理:")

	// 创建带超时的 context（模拟请求超时）
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// 模拟数据库查询
	result := make(chan string, 1)
	go func() {
		// 模拟慢查询
		time.Sleep(200 * time.Millisecond)
		result <- "数据库查询结果"
	}()

	select {
	case data := <-result:
		fmt.Printf("  成功: %s\n", data)
	case <-ctx.Done():
		fmt.Printf("  失败: %v\n", ctx.Err())
	}
}
