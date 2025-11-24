package main

import "fmt"

/*
17 - Defer, Panic 和 Recover

学习目标：
1. 理解 defer 的工作机制
2. 学习 panic 的使用场景
3. 学习 recover 恢复 panic
4. 掌握错误处理的最佳实践

运行方式：
go run 17_defer_panic_recover.go
*/

func main() {
	fmt.Println("=== Defer 基础 ===")

	// defer 延迟执行，在函数返回前执行
	fmt.Println("开始")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("结束")
	// 输出顺序：开始 -> 结束 -> defer 3 -> defer 2 -> defer 1（LIFO）

	fmt.Println("\n=== Defer 执行顺序 ===")

	// defer 按 LIFO（后进先出）顺序执行
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("循环中的 defer: %d\n", i)
	}

	fmt.Println("循环结束")
	// 输出：循环结束 -> 3 -> 2 -> 1

	fmt.Println("\n=== Defer 与参数求值 ===")

	// defer 的参数在 defer 语句执行时求值，而不是在实际执行时
	n := 5
	defer fmt.Printf("defer 时 n 的值: %d\n", n)
	n = 10
	fmt.Printf("当前 n 的值: %d\n", n)
	// 输出：当前 n=10，defer 时 n=5

	fmt.Println("\n=== Defer 与返回值 ===")

	result := deferReturnDemo()
	fmt.Printf("返回值: %d\n", result)

	result2 := namedReturnDemo()
	fmt.Printf("命名返回值: %d\n", result2)

	fmt.Println("\n=== Defer 与资源清理 ===")

	resourceDemo()

	fmt.Println("\n=== Panic 基础 ===")

	// panic 会中断正常执行流程
	// 但 defer 仍会执行
	panicDemo()

	fmt.Println("\n=== Recover 基础 ===")

	// recover 可以捕获 panic
	recoverDemo()

	fmt.Println("\n=== 安全调用函数 ===")

	// 使用 recover 包装可能 panic 的函数
	err := safeCall(func() {
		fmt.Println("执行可能 panic 的函数")
		panic("发生了 panic!")
	})
	if err != nil {
		fmt.Printf("捕获到错误: %v\n", err)
	}

	err = safeCall(func() {
		fmt.Println("执行正常的函数")
	})
	if err == nil {
		fmt.Println("函数正常执行完成")
	}

	fmt.Println("\n=== Defer + Panic + Recover 组合 ===")

	fmt.Println("调用 divideNumbers(10, 2):")
	result3, err := divideNumbers(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %d\n", result3)
	}

	fmt.Println("\n调用 divideNumbers(10, 0):")
	result4, err := divideNumbers(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %d\n", result4)
	}

	fmt.Println("\n=== 多层 Recover ===")

	outerFunc()

	fmt.Println("\n=== Panic 的使用场景 ===")

	fmt.Println(`
Panic 应该在以下场景使用：

1. 不可恢复的错误
   - 程序初始化失败
   - 关键配置缺失
   - 无法恢复的内部错误

2. 检测到不可能发生的情况
   - 表示程序逻辑错误
   - 开发阶段快速失败

3. 初始化时的验证
   - init() 函数中检测配置错误

不应该使用 Panic 的场景：

1. 正常的错误处理
   - 使用 error 返回值
   - 文件不存在、网络错误等

2. 用户输入验证
   - 返回错误信息给用户

3. 可预期的异常情况
   - 应该通过代码逻辑处理
	`)

	fmt.Println("=== Defer 的最佳实践 ===")

	fmt.Println(`
Defer 最佳实践：

1. 资源清理
   f, err := os.Open("file.txt")
   if err != nil { return err }
   defer f.Close()

2. 解锁互斥锁
   mutex.Lock()
   defer mutex.Unlock()

3. 恢复 panic
   defer func() {
       if r := recover(); r != nil {
           log.Printf("Recovered: %v", r)
       }
   }()

4. 记录函数执行时间
   defer func(start time.Time) {
       log.Printf("函数执行时间: %v", time.Since(start))
   }(time.Now())

5. 注意事项
   - defer 有轻微性能开销
   - 避免在循环中使用 defer（除非必要）
   - defer 的参数在声明时求值
	`)

	fmt.Println("\n程序正常结束")
}

// defer 与返回值
func deferReturnDemo() int {
	result := 10
	defer func() {
		result = 20  // 这不会影响返回值
	}()
	return result  // 返回 10
}

// defer 与命名返回值
func namedReturnDemo() (result int) {
	result = 10
	defer func() {
		result = 20  // 这会修改命名返回值
	}()
	return  // 返回 20（命名返回值被 defer 修改）
}

// 资源清理示例
func resourceDemo() {
	fmt.Println("  打开资源 A")
	defer fmt.Println("  关闭资源 A")

	fmt.Println("  打开资源 B")
	defer fmt.Println("  关闭资源 B")

	fmt.Println("  执行操作...")
	// 输出：打开A -> 打开B -> 执行操作 -> 关闭B -> 关闭A
}

// panic 示例
func panicDemo() {
	defer fmt.Println("  panicDemo 的 defer 执行了")

	// 在 goroutine 中捕获 panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  捕获到 panic: %v\n", r)
			}
		}()
		fmt.Println("  即将 panic...")
		panic("这是一个 panic")
		fmt.Println("  这行不会执行")  // 不会执行
	}()

	fmt.Println("  panicDemo 继续执行")
}

// recover 示例
func recoverDemo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  恢复自 panic: %v\n", r)
		}
	}()

	fmt.Println("  recoverDemo 开始")
	panic("测试 panic")
	fmt.Println("  recoverDemo 结束")  // 不会执行
}

// 安全调用函数
func safeCall(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	f()
	return nil
}

// 带有 recover 的除法
func divideNumbers(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	if b == 0 {
		panic("除数不能为零")
	}

	return a / b, nil
}

// 多层函数调用的 panic 传播
func outerFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  outerFunc 捕获: %v\n", r)
		}
	}()

	fmt.Println("  outerFunc 调用 middleFunc")
	middleFunc()
	fmt.Println("  outerFunc 结束")  // 不会执行
}

func middleFunc() {
	defer fmt.Println("  middleFunc 的 defer")
	fmt.Println("  middleFunc 调用 innerFunc")
	innerFunc()
	fmt.Println("  middleFunc 结束")  // 不会执行
}

func innerFunc() {
	defer fmt.Println("  innerFunc 的 defer")
	fmt.Println("  innerFunc 开始")
	panic("来自 innerFunc 的 panic")
	fmt.Println("  innerFunc 结束")  // 不会执行
}
