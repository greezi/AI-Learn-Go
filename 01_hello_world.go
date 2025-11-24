package main

import "fmt"

/*
01 - Hello World 与基础程序结构

学习目标：
1. 了解 Go 程序的基本结构
2. 理解 package 和 import 的概念
3. 学习如何打印输出

运行方式：
go run 01_hello_world.go
*/

func main() {
	// main 是程序的入口函数
	// 每个可执行的 Go 程序都必须有一个 main 包和 main 函数

	// 使用 fmt.Println 打印一行文本
	fmt.Println("Hello, World!")

	// fmt.Print 不会自动换行
	fmt.Print("Hello, ")
	fmt.Print("Go!\n")

	// fmt.Printf 支持格式化输出
	name := "Golang"
	version := 1.21
	fmt.Printf("欢迎学习 %s，版本 %.2f\n", name, version)

	// 多行注释示例
	/*
		这是一个多行注释
		可以写很多行
	*/
}
