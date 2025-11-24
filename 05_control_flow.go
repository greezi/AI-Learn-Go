package main

import "fmt"

/*
05 - 控制流程：if 和 switch

学习目标：
1. 学习 if-else 条件语句
2. 学习 switch 语句
3. 理解 Go 特有的语法特性

运行方式：
go run 05_control_flow.go
*/

func main() {
	fmt.Println("=== if 语句 ===")

	age := 18
	if age >= 18 {
		fmt.Println("已成年")
	}

	// if-else
	score := 85
	if score >= 90 {
		fmt.Println("成绩优秀")
	} else if score >= 60 {
		fmt.Println("成绩及格")
	} else {
		fmt.Println("成绩不及格")
	}

	// if 语句可以包含初始化语句
	// 变量的作用域仅在 if 块内
	if num := 10; num%2 == 0 {
		fmt.Printf("%d 是偶数\n", num)
	} else {
		fmt.Printf("%d 是奇数\n", num)
	}
	// num 在这里不可访问

	fmt.Println("\n=== switch 语句 ===")

	// 基本 switch
	day := 3
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6, 7:  // 多个条件
		fmt.Println("周末")
	default:
		fmt.Println("无效的日期")
	}

	// switch 带初始化语句
	switch hour := 14; {
	case hour < 12:
		fmt.Println("上午")
	case hour < 18:
		fmt.Println("下午")
	default:
		fmt.Println("晚上")
	}

	// 无条件 switch（相当于 if-else 链）
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("非常冷")
	case temperature < 15:
		fmt.Println("冷")
	case temperature < 25:
		fmt.Println("温暖")
	default:
		fmt.Println("热")
	}

	// switch 匹配类型
	var value interface{} = "hello"
	switch v := value.(type) {
	case int:
		fmt.Printf("整数: %d\n", v)
	case string:
		fmt.Printf("字符串: %s\n", v)
	case bool:
		fmt.Printf("布尔: %v\n", v)
	default:
		fmt.Printf("未知类型: %T\n", v)
	}

	// fallthrough 关键字（穿透到下一个 case）
	number := 2
	fmt.Println("\nfallthrough 示例:")
	switch number {
	case 1:
		fmt.Println("数字是 1")
	case 2:
		fmt.Println("数字是 2")
		fallthrough  // 继续执行下一个 case
	case 3:
		fmt.Println("数字是 2 或 3")
	default:
		fmt.Println("其他数字")
	}

	fmt.Println("\n=== 复杂条件判断 ===")

	// 嵌套 if
	x := 10
	y := 20
	if x > 0 {
		if y > 0 {
			fmt.Println("x 和 y 都是正数")
		}
	}

	// 使用逻辑运算符简化
	if x > 0 && y > 0 {
		fmt.Println("x 和 y 都是正数（简化版）")
	}

	// 嵌套 switch
	category := "电子产品"
	subcategory := "手机"
	switch category {
	case "电子产品":
		fmt.Println("类别：电子产品")
		switch subcategory {
		case "电脑":
			fmt.Println("  子类别：电脑")
		case "手机":
			fmt.Println("  子类别：手机")
		}
	case "图书":
		fmt.Println("类别：图书")
	}
}
