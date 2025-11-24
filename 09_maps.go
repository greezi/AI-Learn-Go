package main

import "fmt"

/*
09 - 映射（Map）

学习目标：
1. 学习 map 的声明和初始化
2. 学习 map 的增删改查操作
3. 理解 map 的特性
4. 掌握 map 的遍历

运行方式：
go run 09_maps.go
*/

func main() {
	fmt.Println("=== Map 声明和初始化 ===")

	// 声明 map（未初始化，值为 nil）
	var map1 map[string]int
	fmt.Printf("map1: %v, 是否为 nil: %v\n", map1, map1 == nil)
	// 注意：不能向 nil map 添加元素，会 panic

	// 使用 make 创建 map
	map2 := make(map[string]int)
	fmt.Printf("map2: %v, 是否为 nil: %v\n", map2, map2 == nil)

	// 使用字面量初始化
	map3 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 7,
	}
	fmt.Printf("map3: %v\n", map3)

	// 不同类型的 map
	var (
		intToString    = map[int]string{1: "一", 2: "二", 3: "三"}
		stringToBool   = map[string]bool{"active": true, "deleted": false}
		stringToSlice  = map[string][]int{"odds": {1, 3, 5}, "evens": {2, 4, 6}}
	)
	fmt.Printf("intToString: %v\n", intToString)
	fmt.Printf("stringToBool: %v\n", stringToBool)
	fmt.Printf("stringToSlice: %v\n", stringToSlice)

	fmt.Println("\n=== Map 基本操作 ===")

	// 创建一个学生成绩 map
	scores := make(map[string]int)

	// 添加/修改元素
	scores["张三"] = 85
	scores["李四"] = 92
	scores["王五"] = 78
	fmt.Printf("成绩表: %v\n", scores)

	// 访问元素
	fmt.Printf("张三的成绩: %d\n", scores["张三"])

	// 访问不存在的键（返回零值）
	fmt.Printf("赵六的成绩: %d (不存在，返回零值)\n", scores["赵六"])

	// 检查键是否存在（推荐方式）
	score, exists := scores["张三"]
	if exists {
		fmt.Printf("张三的成绩: %d\n", score)
	} else {
		fmt.Println("张三不存在")
	}

	score2, exists2 := scores["赵六"]
	if exists2 {
		fmt.Printf("赵六的成绩: %d\n", score2)
	} else {
		fmt.Println("赵六不存在")
	}

	// 修改元素
	scores["张三"] = 90
	fmt.Printf("修改后张三的成绩: %d\n", scores["张三"])

	// 删除元素
	delete(scores, "王五")
	fmt.Printf("删除王五后: %v\n", scores)

	// 删除不存在的键（不会出错）
	delete(scores, "不存在的人")

	// 获取 map 的长度
	fmt.Printf("成绩表人数: %d\n", len(scores))

	fmt.Println("\n=== Map 遍历 ===")

	students := map[string]int{
		"张三": 85,
		"李四": 92,
		"王五": 78,
		"赵六": 88,
	}

	// 遍历键值对
	fmt.Println("所有学生成绩:")
	for name, score := range students {
		fmt.Printf("  %s: %d 分\n", name, score)
	}

	// 只遍历键
	fmt.Println("所有学生姓名:")
	for name := range students {
		fmt.Printf("  %s\n", name)
	}

	// 只遍历值（通过忽略键）
	fmt.Println("所有成绩:")
	for _, score := range students {
		fmt.Printf("  %d 分\n", score)
	}

	// 注意：map 的遍历顺序是随机的
	fmt.Println("\n多次遍历可能得到不同顺序:")
	for i := 0; i < 3; i++ {
		fmt.Printf("第 %d 次: ", i+1)
		for name := range students {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}

	fmt.Println("\n=== Map 嵌套 ===")

	// map 的值可以是另一个 map
	userInfo := map[string]map[string]string{
		"user1": {
			"name":  "张三",
			"email": "zhangsan@example.com",
			"city":  "北京",
		},
		"user2": {
			"name":  "李四",
			"email": "lisi@example.com",
			"city":  "上海",
		},
	}

	fmt.Println("用户信息:")
	for userID, info := range userInfo {
		fmt.Printf("  %s:\n", userID)
		for key, value := range info {
			fmt.Printf("    %s: %s\n", key, value)
		}
	}

	// 访问嵌套 map
	fmt.Printf("user1 的邮箱: %s\n", userInfo["user1"]["email"])

	// 修改嵌套 map
	userInfo["user1"]["city"] = "深圳"
	fmt.Printf("修改后 user1 的城市: %s\n", userInfo["user1"]["city"])

	fmt.Println("\n=== Map 作为集合（Set） ===")

	// Go 没有内置的 Set，但可以用 map[type]bool 实现
	set := make(map[string]bool)

	// 添加元素
	set["apple"] = true
	set["banana"] = true
	set["orange"] = true
	fmt.Printf("集合: %v\n", set)

	// 检查元素是否存在
	if set["apple"] {
		fmt.Println("apple 在集合中")
	}

	if !set["grape"] {
		fmt.Println("grape 不在集合中")
	}

	// 删除元素
	delete(set, "banana")

	// 遍历集合
	fmt.Print("集合元素: ")
	for item := range set {
		fmt.Printf("%s ", item)
	}
	fmt.Println()

	// 使用 struct{} 作为值类型更节省内存
	efficientSet := make(map[string]struct{})
	efficientSet["item1"] = struct{}{}
	efficientSet["item2"] = struct{}{}

	_, exists = efficientSet["item1"]
	fmt.Printf("item1 在集合中: %v\n", exists)

	fmt.Println("\n=== 实用示例 ===")

	// 示例1：统计单词出现次数
	wordCount := make(map[string]int)
	words := []string{"hello", "world", "hello", "go", "go", "go"}

	for _, word := range words {
		wordCount[word]++
	}
	fmt.Println("单词计数:")
	for word, count := range wordCount {
		fmt.Printf("  %s: %d 次\n", word, count)
	}

	// 示例2：分组
	ages := map[string]int{
		"张三": 25,
		"李四": 30,
		"王五": 25,
		"赵六": 30,
		"钱七": 35,
	}

	ageGroups := make(map[int][]string)
	for name, age := range ages {
		ageGroups[age] = append(ageGroups[age], name)
	}

	fmt.Println("\n年龄分组:")
	for age, names := range ageGroups {
		fmt.Printf("  %d 岁: %v\n", age, names)
	}

	// 示例3：反转 map（键值互换）
	original := map[string]int{"a": 1, "b": 2, "c": 3}
	reversed := make(map[int]string)
	for key, value := range original {
		reversed[value] = key
	}
	fmt.Printf("\n原始 map: %v\n", original)
	fmt.Printf("反转后: %v\n", reversed)

	// 示例4：合并两个 map
	mapA := map[string]int{"a": 1, "b": 2}
	mapB := map[string]int{"c": 3, "d": 4}
	merged := make(map[string]int)

	for k, v := range mapA {
		merged[k] = v
	}
	for k, v := range mapB {
		merged[k] = v
	}
	fmt.Printf("\n合并后的 map: %v\n", merged)
}
