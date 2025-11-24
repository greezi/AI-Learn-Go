package main

import "fmt"

/*
08 - 数组和切片

学习目标：
1. 学习数组的声明和使用
2. 学习切片的概念和操作
3. 理解数组和切片的区别
4. 掌握切片的常用操作

运行方式：
go run 08_arrays_slices.go
*/

func main() {
	fmt.Println("=== 数组 ===")

	// 声明并初始化数组
	var arr1 [5]int  // 声明一个长度为 5 的整数数组，默认值都是 0
	fmt.Printf("arr1: %v\n", arr1)

	// 初始化数组
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr2: %v\n", arr2)

	// 部分初始化
	arr3 := [5]int{1, 2, 3}  // 后面的元素为 0
	fmt.Printf("arr3: %v\n", arr3)

	// 让编译器推断数组长度
	arr4 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr4: %v, 长度: %d\n", arr4, len(arr4))

	// 指定索引初始化
	arr5 := [5]int{0: 10, 2: 30, 4: 50}
	fmt.Printf("arr5: %v\n", arr5)

	// 访问和修改数组元素
	fmt.Printf("arr2[0] = %d\n", arr2[0])
	arr2[0] = 100
	fmt.Printf("修改后 arr2[0] = %d\n", arr2[0])

	// 遍历数组
	fmt.Println("遍历 arr2:")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("  arr2[%d] = %d\n", i, arr2[i])
	}

	// 使用 range 遍历
	fmt.Println("使用 range 遍历:")
	for index, value := range arr2 {
		fmt.Printf("  索引 %d: 值 %d\n", index, value)
	}

	// 数组是值类型（复制传递）
	arr6 := arr2  // 复制整个数组
	arr6[0] = 999
	fmt.Printf("arr2[0] = %d, arr6[0] = %d (互不影响)\n", arr2[0], arr6[0])

	// 多维数组
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("3x3 矩阵:\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	fmt.Println("\n=== 切片 ===")

	// 切片是对数组的引用，更灵活
	var slice1 []int  // 声明切片（未初始化，值为 nil）
	fmt.Printf("slice1: %v, 长度: %d, 容量: %d, 是否为 nil: %v\n",
		slice1, len(slice1), cap(slice1), slice1 == nil)

	// 使用字面量创建切片
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice2: %v, 长度: %d, 容量: %d\n",
		slice2, len(slice2), cap(slice2))

	// 使用 make 创建切片
	slice3 := make([]int, 5)      // 长度和容量都是 5
	slice4 := make([]int, 3, 10)  // 长度 3，容量 10
	fmt.Printf("slice3: %v, 长度: %d, 容量: %d\n",
		slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4: %v, 长度: %d, 容量: %d\n",
		slice4, len(slice4), cap(slice4))

	// 从数组创建切片
	arr := [5]int{10, 20, 30, 40, 50}
	slice5 := arr[1:4]  // 包含索引 1, 2, 3（不包含 4）
	fmt.Printf("slice5 (arr[1:4]): %v\n", slice5)

	slice6 := arr[:3]   // 从开始到索引 3（不包含）
	slice7 := arr[2:]   // 从索引 2 到结束
	slice8 := arr[:]    // 整个数组
	fmt.Printf("slice6 (arr[:3]): %v\n", slice6)
	fmt.Printf("slice7 (arr[2:]): %v\n", slice7)
	fmt.Printf("slice8 (arr[:]): %v\n", slice8)

	fmt.Println("\n=== 切片操作 ===")

	// append - 追加元素
	nums := []int{1, 2, 3}
	fmt.Printf("初始: %v\n", nums)
	nums = append(nums, 4)
	fmt.Printf("追加 4: %v\n", nums)
	nums = append(nums, 5, 6, 7)
	fmt.Printf("追加 5,6,7: %v\n", nums)

	// 追加另一个切片
	nums2 := []int{8, 9, 10}
	nums = append(nums, nums2...)
	fmt.Printf("追加另一个切片: %v\n", nums)

	// copy - 复制切片
	source := []int{1, 2, 3, 4, 5}
	dest := make([]int, 5)
	count := copy(dest, source)
	fmt.Printf("复制了 %d 个元素: %v\n", count, dest)

	// 部分复制
	dest2 := make([]int, 3)
	copy(dest2, source)
	fmt.Printf("部分复制: %v\n", dest2)

	// 删除元素（通过切片操作）
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("原始: %v\n", slice)

	// 删除索引 2 的元素
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Printf("删除索引 2: %v\n", slice)

	// 插入元素
	slice = []int{1, 2, 4, 5}
	fmt.Printf("原始: %v\n", slice)

	// 在索引 2 插入 3
	index = 2
	value := 3
	slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
	fmt.Printf("在索引 2 插入 3: %v\n", slice)

	fmt.Println("\n=== 切片是引用类型 ===")

	// 切片是引用，修改会影响原始数据
	original := []int{1, 2, 3, 4, 5}
	reference := original
	reference[0] = 999
	fmt.Printf("original: %v, reference: %v (共享底层数组)\n", original, reference)

	// 使用 copy 创建独立副本
	original2 := []int{1, 2, 3, 4, 5}
	independent := make([]int, len(original2))
	copy(independent, original2)
	independent[0] = 999
	fmt.Printf("original2: %v, independent: %v (独立副本)\n", original2, independent)

	fmt.Println("\n=== 二维切片 ===")

	// 创建二维切片
	matrix2D := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("二维切片:")
	for i, row := range matrix2D {
		fmt.Printf("  行 %d: %v\n", i, row)
	}

	// 动态创建二维切片
	rows, cols := 3, 4
	dynamic2D := make([][]int, rows)
	for i := range dynamic2D {
		dynamic2D[i] = make([]int, cols)
	}
	fmt.Printf("动态二维切片 (%dx%d): %v\n", rows, cols, dynamic2D)

	fmt.Println("\n=== 实用示例 ===")

	// 过滤切片
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evens []int
	for _, n := range numbers {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	fmt.Printf("偶数: %v\n", evens)

	// 反转切片
	toReverse := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(toReverse)-1; i < j; i, j = i+1, j-1 {
		toReverse[i], toReverse[j] = toReverse[j], toReverse[i]
	}
	fmt.Printf("反转后: %v\n", toReverse)
}
