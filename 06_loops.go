package main

import "fmt"

/*
06 - 循环

学习目标：
1. 学习 for 循环（Go 中唯一的循环结构）
2. 学习 for 的各种形式
3. 学习 break、continue 和 goto

运行方式：
go run 06_loops.go
*/

func main() {
	fmt.Println("=== 基本 for 循环 ===")

	// 标准 for 循环（类似 C/Java）
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 省略初始化语句
	j := 1
	for ; j <= 5; j++ {
		fmt.Printf("%d ", j)
	}
	fmt.Println()

	fmt.Println("\n=== while 风格的 for 循环 ===")

	// Go 没有 while，但可以用 for 实现
	count := 0
	for count < 5 {
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println()

	fmt.Println("\n=== 无限循环 ===")

	// 无限循环（需要用 break 退出）
	num := 0
	for {
		num++
		if num > 3 {
			break  // 跳出循环
		}
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	fmt.Println("\n=== range 遍历 ===")

	// 遍历数组/切片
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("遍历切片（索引和值）:")
	for index, value := range numbers {
		fmt.Printf("  索引 %d: 值 %d\n", index, value)
	}

	// 只要值，忽略索引
	fmt.Println("只要值:")
	for _, value := range numbers {
		fmt.Printf("  %d ", value)
	}
	fmt.Println()

	// 只要索引
	fmt.Println("只要索引:")
	for index := range numbers {
		fmt.Printf("  %d ", index)
	}
	fmt.Println()

	// 遍历字符串
	str := "Hello,世界"
	fmt.Println("\n遍历字符串:")
	for index, char := range str {
		fmt.Printf("  位置 %d: %c (Unicode: %U)\n", index, char, char)
	}

	// 遍历 map
	scores := map[string]int{
		"张三": 85,
		"李四": 92,
		"王五": 78,
	}
	fmt.Println("\n遍历 map:")
	for name, score := range scores {
		fmt.Printf("  %s: %d 分\n", name, score)
	}

	fmt.Println("\n=== continue 语句 ===")

	// continue 跳过本次循环
	fmt.Println("打印奇数:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue  // 跳过偶数
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Println("\n=== break 语句 ===")

	// break 终止循环
	fmt.Println("找到第一个大于 30 的数:")
	nums := []int{10, 20, 35, 40, 50}
	for _, n := range nums {
		if n > 30 {
			fmt.Printf("找到了: %d\n", n)
			break
		}
	}

	fmt.Println("\n=== 标签和 goto ===")

	// 使用标签跳出多层循环
	fmt.Println("标签示例（跳出嵌套循环）:")
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d,%d) ", i, j)
			if i == 2 && j == 2 {
				break OuterLoop  // 跳出外层循环
			}
		}
		fmt.Println()
	}
	fmt.Println("已跳出")

	// goto 语句（不推荐过度使用）
	fmt.Println("\ngoto 示例:")
	x := 0
Loop:
	x++
	fmt.Printf("%d ", x)
	if x < 5 {
		goto Loop
	}
	fmt.Println()

	fmt.Println("\n=== 嵌套循环 ===")

	// 打印乘法表
	fmt.Println("九九乘法表:")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%2d ", j, i, i*j)
		}
		fmt.Println()
	}

	fmt.Println("\n=== 实用示例 ===")

	// 示例1：求和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("1 到 100 的和: %d\n", sum)

	// 示例2：阶乘
	factorial := 1
	n := 5
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Printf("%d 的阶乘: %d\n", n, factorial)

	// 示例3：斐波那契数列
	fmt.Printf("斐波那契数列前 10 项: ")
	a, b := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}
