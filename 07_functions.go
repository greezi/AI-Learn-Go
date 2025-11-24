package main

import "fmt"

/*
07 - 函数

学习目标：
1. 学习函数的定义和调用
2. 学习参数和返回值
3. 学习多返回值
4. 学习可变参数
5. 学习匿名函数和闭包

运行方式：
go run 07_functions.go
*/

func main() {
	fmt.Println("=== 基本函数 ===")

	// 调用无参数无返回值的函数
	sayHello()

	// 调用有参数的函数
	sayHelloTo("张三")

	// 调用有返回值的函数
	result := addTwoNumbers(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)

	fmt.Println("\n=== 多返回值 ===")

	// Go 函数可以返回多个值
	sumResult, diffResult := calcSumAndDiff(10, 3)
	fmt.Printf("和: %d, 差: %d\n", sumResult, diffResult)

	// 忽略某些返回值
	sumResult2, _ := calcSumAndDiff(15, 5)
	fmt.Printf("只要和: %d\n", sumResult2)

	// 函数返回多个值的常见用法：返回值和错误
	value, err := divideFloat(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %f\n", value)
	}

	value2, err2 := divideFloat(10, 0)
	if err2 != nil {
		fmt.Printf("错误: %v\n", err2)
	} else {
		fmt.Printf("结果: %f\n", value2)
	}

	fmt.Println("\n=== 命名返回值 ===")

	// 返回值可以命名
	quotient, remainder := divideWithRemainder(17, 5)
	fmt.Printf("17 ÷ 5 = %d ... %d\n", quotient, remainder)

	fmt.Println("\n=== 可变参数 ===")

	// 函数可以接受可变数量的参数
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Printf("1+2+3+4+5 = %d\n", total)

	// 传递切片
	nums := []int{10, 20, 30}
	total2 := sumAll(nums...)  // 使用 ... 展开切片
	fmt.Printf("10+20+30 = %d\n", total2)

	// 混合参数
	printDetails("用户信息", "张三", "25岁", "北京")

	fmt.Println("\n=== 函数作为值 ===")

	// 函数可以赋值给变量
	var f func(int, int) int
	f = addTwoNumbers
	fmt.Printf("使用函数变量: %d\n", f(5, 7))

	// 函数作为参数
	runOperation(10, 5, addTwoNumbers)
	runOperation(10, 5, multiplyTwoNumbers)

	fmt.Println("\n=== 匿名函数 ===")

	// 定义并立即调用匿名函数
	func() {
		fmt.Println("这是一个匿名函数")
	}()

	// 匿名函数赋值给变量
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("5 的平方: %d\n", square(5))

	fmt.Println("\n=== 闭包 ===")

	// 闭包：函数可以捕获外部变量
	counter := createCounter()
	fmt.Printf("计数: %d\n", counter())  // 1
	fmt.Printf("计数: %d\n", counter())  // 2
	fmt.Printf("计数: %d\n", counter())  // 3

	// 每个闭包都有自己的状态
	counter2 := createCounter()
	fmt.Printf("新计数器: %d\n", counter2())  // 1

	// 闭包的实用例子：累加器
	adder := createAdder(10)
	fmt.Printf("10 + 5 = %d\n", adder(5))
	fmt.Printf("10 + 20 = %d\n", adder(20))

	fmt.Println("\n=== 递归函数 ===")

	// 计算阶乘
	fmt.Printf("5! = %d\n", calcFactorial(5))

	// 斐波那契数列
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", calcFibonacci(i))
	}
	fmt.Println()

	fmt.Println("\n=== 延迟执行（defer） ===")

	// defer 会在函数返回前执行
	showDeferExample()
}

// 无参数无返回值的函数
func sayHello() {
	fmt.Println("你好，Go！")
}

// 有参数的函数
func sayHelloTo(name string) {
	fmt.Printf("你好，%s！\n", name)
}

// 有返回值的函数
func addTwoNumbers(a, b int) int {
	return a + b
}

// 多返回值
func calcSumAndDiff(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// 返回值和错误
func divideFloat(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	return a / b, nil
}

// 命名返回值
func divideWithRemainder(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return  // 可以省略返回值名称
}

// 可变参数
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 混合普通参数和可变参数
func printDetails(title string, details ...string) {
	fmt.Printf("%s:\n", title)
	for _, item := range details {
		fmt.Printf("  - %s\n", item)
	}
}

// 用于演示函数作为参数
func multiplyTwoNumbers(a, b int) int {
	return a * b
}

func runOperation(a, b int, op func(int, int) int) {
	result := op(a, b)
	fmt.Printf("操作结果: %d\n", result)
}

// 返回闭包的函数
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func createAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// 递归函数：阶乘
func calcFactorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * calcFactorial(n-1)
}

// 递归函数：斐波那契
func calcFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return calcFibonacci(n-1) + calcFibonacci(n-2)
}

// defer 示例
func showDeferExample() {
	fmt.Println("函数开始")
	defer fmt.Println("defer 1: 这会最后执行")
	defer fmt.Println("defer 2: 这会倒数第二执行")
	fmt.Println("函数中间")
	defer fmt.Println("defer 3: 这会倒数第三执行")
	fmt.Println("函数结束")
	// defer 按照 LIFO（后进先出）顺序执行
}
