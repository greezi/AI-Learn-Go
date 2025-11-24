package main

import "fmt"

/*
02 - 变量和数据类型

学习目标：
1. 学习变量声明的多种方式
2. 了解 Go 的基本数据类型
3. 理解类型推断和零值

运行方式：
go run 02_variables_types.go
*/

func main() {
	fmt.Println("=== 变量声明 ===")

	// 方式1：使用 var 关键字声明并初始化
	var name string = "张三"
	fmt.Println("姓名:", name)

	// 方式2：类型推断（省略类型）
	var age = 25
	fmt.Println("年龄:", age)

	// 方式3：短变量声明（最常用，只能在函数内使用）
	city := "北京"
	fmt.Println("城市:", city)

	// 方式4：同时声明多个变量
	var x, y, z int = 1, 2, 3
	fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)

	// 方式5：变量组
	var (
		username = "user123"
		password = "pass456"
		isAdmin  = false
	)
	fmt.Printf("用户: %s, 密码: %s, 管理员: %v\n", username, password, isAdmin)

	fmt.Println("\n=== 基本数据类型 ===")

	// 布尔型
	var isActive bool = true
	fmt.Printf("布尔型: %v (类型: %T)\n", isActive, isActive)

	// 字符串
	var message string = "你好，Go！"
	fmt.Printf("字符串: %s (类型: %T)\n", message, message)

	// 整数类型
	var (
		num8   int8   = 127           // 8位整数 (-128 到 127)
		num16  int16  = 32767         // 16位整数
		num32  int32  = 2147483647    // 32位整数
		num64  int64  = 9223372036854775807 // 64位整数
		numInt int    = 100           // 根据平台自动选择 32 或 64 位
	)
	fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d, int: %d\n",
		num8, num16, num32, num64, numInt)

	// 无符号整数
	var (
		unum8  uint8  = 255
		unum16 uint16 = 65535
		unum32 uint32 = 4294967295
	)
	fmt.Printf("uint8: %d, uint16: %d, uint32: %d\n", unum8, unum16, unum32)

	// 浮点数
	var (
		float32Num float32 = 3.14
		float64Num float64 = 3.141592653589793
	)
	fmt.Printf("float32: %f, float64: %.15f\n", float32Num, float64Num)

	// 复数
	var complexNum complex64 = 1 + 2i
	fmt.Printf("复数: %v\n", complexNum)

	// byte (uint8 的别名) 和 rune (int32 的别名，用于 Unicode)
	var (
		byteVar byte = 'A'  // ASCII 字符
		runeVar rune = '中' // Unicode 字符
	)
	fmt.Printf("byte: %c (%d), rune: %c (%d)\n", byteVar, byteVar, runeVar, runeVar)

	fmt.Println("\n=== 零值 ===")
	// 未初始化的变量会被赋予零值
	var (
		zeroInt    int
		zeroFloat  float64
		zeroBool   bool
		zeroString string
	)
	fmt.Printf("int 零值: %d\n", zeroInt)
	fmt.Printf("float64 零值: %f\n", zeroFloat)
	fmt.Printf("bool 零值: %v\n", zeroBool)
	fmt.Printf("string 零值: '%s' (空字符串)\n", zeroString)

	fmt.Println("\n=== 类型转换 ===")
	// Go 需要显式类型转换
	var intNum int = 42
	var floatNum float64 = float64(intNum)
	var uintNum uint = uint(intNum)
	fmt.Printf("int: %d -> float64: %f, uint: %d\n", intNum, floatNum, uintNum)
}
