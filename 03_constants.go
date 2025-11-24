package main

import "fmt"

/*
03 - 常量

学习目标：
1. 学习常量的声明和使用
2. 了解 iota 枚举器
3. 理解常量的类型

运行方式：
go run 03_constants.go
*/

func main() {
	fmt.Println("=== 常量声明 ===")

	// 使用 const 关键字声明常量
	const pi = 3.14159
	const greeting = "你好"
	fmt.Printf("圆周率: %f\n", pi)
	fmt.Printf("问候语: %s\n", greeting)

	// 常量组
	const (
		StatusOK      = 200
		StatusCreated = 201
		StatusBadRequest = 400
		StatusNotFound = 404
	)
	fmt.Printf("HTTP 状态码 - OK: %d, Not Found: %d\n", StatusOK, StatusNotFound)

	// 类型化常量
	const typedInt int = 100
	const typedString string = "类型化字符串"
	fmt.Printf("类型化常量: %d (%T), %s (%T)\n", typedInt, typedInt, typedString, typedString)

	fmt.Println("\n=== iota 枚举器 ===")

	// iota 在 const 中用于创建枚举值
	// iota 从 0 开始，每行递增 1
	const (
		Sunday = iota    // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Printf("星期日: %d, 星期一: %d, 星期六: %d\n", Sunday, Monday, Saturday)

	// iota 可以参与表达式
	const (
		_  = iota             // 0 (使用 _ 忽略)
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB                    // 1 << 20 = 1048576
		GB                    // 1 << 30 = 1073741824
		TB                    // 1 << 40
	)
	fmt.Printf("1 KB = %d 字节\n", KB)
	fmt.Printf("1 MB = %d 字节\n", MB)
	fmt.Printf("1 GB = %d 字节\n", GB)

	// 多个 iota 在同一行
	const (
		a, b = iota + 1, iota + 2  // 1, 2
		c, d                        // 2, 3
		e, f                        // 3, 4
	)
	fmt.Printf("a=%d, b=%d, c=%d, d=%d, e=%d, f=%d\n", a, b, c, d, e, f)

	// iota 跳过某些值
	const (
		n1 = iota // 0
		n2        // 1
		_         // 2 (跳过)
		n4        // 3
	)
	fmt.Printf("n1=%d, n2=%d, n4=%d\n", n1, n2, n4)

	fmt.Println("\n=== 常量的特性 ===")

	// 常量可以是无类型的，可以被赋值给不同类型的变量
	const untypedConst = 42
	var intVar int = untypedConst
	var floatVar float64 = untypedConst
	var complexVar complex128 = untypedConst
	fmt.Printf("无类型常量可赋值给: int=%d, float64=%f, complex128=%v\n",
		intVar, floatVar, complexVar)

	// 注意：常量不能被修改
	// pi = 3.14  // 这行会导致编译错误：cannot assign to pi
}
