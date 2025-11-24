package main

import "fmt"

/*
04 - 运算符

学习目标：
1. 学习算术运算符
2. 学习比较运算符
3. 学习逻辑运算符
4. 学习位运算符
5. 学习赋值运算符

运行方式：
go run 04_operators.go
*/

func main() {
	fmt.Println("=== 算术运算符 ===")

	a, b := 10, 3
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("加法: a + b = %d\n", a+b)
	fmt.Printf("减法: a - b = %d\n", a-b)
	fmt.Printf("乘法: a * b = %d\n", a*b)
	fmt.Printf("除法: a / b = %d\n", a/b)
	fmt.Printf("取模: a %% b = %d\n", a%b)

	// 自增和自减（只有后置，没有前置）
	count := 5
	count++  // count = count + 1
	fmt.Printf("count++ = %d\n", count)
	count--  // count = count - 1
	fmt.Printf("count-- = %d\n", count)

	fmt.Println("\n=== 比较运算符 ===")

	x, y := 10, 20
	fmt.Printf("x = %d, y = %d\n", x, y)
	fmt.Printf("x == y: %v\n", x == y) // 等于
	fmt.Printf("x != y: %v\n", x != y) // 不等于
	fmt.Printf("x < y: %v\n", x < y)   // 小于
	fmt.Printf("x <= y: %v\n", x <= y) // 小于等于
	fmt.Printf("x > y: %v\n", x > y)   // 大于
	fmt.Printf("x >= y: %v\n", x >= y) // 大于等于

	fmt.Println("\n=== 逻辑运算符 ===")

	p, q := true, false
	fmt.Printf("p = %v, q = %v\n", p, q)
	fmt.Printf("p && q (逻辑与): %v\n", p && q)
	fmt.Printf("p || q (逻辑或): %v\n", p || q)
	fmt.Printf("!p (逻辑非): %v\n", !p)

	// 短路求值
	fmt.Println("\n短路求值示例:")
	result := (x > 5) && (y < 30)
	fmt.Printf("(x > 5) && (y < 30) = %v\n", result)

	fmt.Println("\n=== 位运算符 ===")

	m, n := 12, 25  // 二进制: 1100, 11001
	fmt.Printf("m = %d (二进制: %b)\n", m, m)
	fmt.Printf("n = %d (二进制: %b)\n", n, n)
	fmt.Printf("m & n (按位与): %d (二进制: %b)\n", m&n, m&n)
	fmt.Printf("m | n (按位或): %d (二进制: %b)\n", m|n, m|n)
	fmt.Printf("m ^ n (按位异或): %d (二进制: %b)\n", m^n, m^n)
	fmt.Printf("^m (按位取反): %d\n", ^m)

	// 位移运算
	num := 8  // 二进制: 1000
	fmt.Printf("\nnum = %d (二进制: %b)\n", num, num)
	fmt.Printf("num << 2 (左移): %d (二进制: %b)\n", num<<2, num<<2)
	fmt.Printf("num >> 2 (右移): %d (二进制: %b)\n", num>>2, num>>2)

	fmt.Println("\n=== 赋值运算符 ===")

	value := 10
	fmt.Printf("初始值: %d\n", value)

	value += 5  // value = value + 5
	fmt.Printf("value += 5: %d\n", value)

	value -= 3  // value = value - 3
	fmt.Printf("value -= 3: %d\n", value)

	value *= 2  // value = value * 2
	fmt.Printf("value *= 2: %d\n", value)

	value /= 4  // value = value / 4
	fmt.Printf("value /= 4: %d\n", value)

	value %= 3  // value = value % 3
	fmt.Printf("value %%= 3: %d\n", value)

	// 位运算赋值
	bits := 12
	fmt.Printf("\n初始值: %d (二进制: %b)\n", bits, bits)
	bits &= 10  // bits = bits & 10
	fmt.Printf("bits &= 10: %d (二进制: %b)\n", bits, bits)
	bits |= 5   // bits = bits | 5
	fmt.Printf("bits |= 5: %d (二进制: %b)\n", bits, bits)
	bits ^= 3   // bits = bits ^ 3
	fmt.Printf("bits ^= 3: %d (二进制: %b)\n", bits, bits)
	bits <<= 1  // bits = bits << 1
	fmt.Printf("bits <<= 1: %d (二进制: %b)\n", bits, bits)
	bits >>= 1  // bits = bits >> 1
	fmt.Printf("bits >>= 1: %d (二进制: %b)\n", bits, bits)

	fmt.Println("\n=== 其他运算符 ===")

	// 取地址运算符 &
	number := 42
	ptr := &number
	fmt.Printf("number 的值: %d\n", number)
	fmt.Printf("number 的地址: %p\n", ptr)

	// 取值运算符 *
	fmt.Printf("ptr 指向的值: %d\n", *ptr)
}
