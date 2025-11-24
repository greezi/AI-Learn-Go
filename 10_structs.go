package main

import "fmt"

/*
10 - 结构体

学习目标：
1. 学习结构体的定义和初始化
2. 学习结构体字段的访问和修改
3. 学习匿名字段和嵌入
4. 理解结构体的标签

运行方式：
go run 10_structs.go
*/

// 定义一个简单的结构体
type PersonInfo struct {
	Name string
	Age  int
	City string
}

// 带有不同类型字段的结构体
type Staff struct {
	ID       int
	Name     string
	Position string
	Salary   float64
	IsActive bool
}

// 嵌套结构体
type Location struct {
	Street  string
	City    string
	ZipCode string
}

type Pupil struct {
	Name     string
	Age      int
	Location Location  // 嵌套结构体
	Scores   []int     // 切片字段
}

// 匿名字段（嵌入）
type Contact struct {
	Email string
	Phone string
}

type Member struct {
	Name    string
	Contact  // 匿名字段，字段名与类型名相同
}

// 结构体标签（用于 JSON、数据库等）
type Goods struct {
	ID    int     `json:"id"`
	Name  string  `json:"product_name"`
	Price float64 `json:"price,omitempty"`
}

// 空结构体（不占用内存）
type Placeholder struct{}

func main() {
	fmt.Println("=== 结构体基础 ===")

	// 方式1：声明并初始化（字段为零值）
	var p1 PersonInfo
	fmt.Printf("p1: %+v\n", p1)  // %+v 打印字段名和值

	// 方式2：使用字面量初始化（按顺序）
	p2 := PersonInfo{"张三", 25, "北京"}
	fmt.Printf("p2: %+v\n", p2)

	// 方式3：使用字段名初始化（推荐，更清晰）
	p3 := PersonInfo{
		Name: "李四",
		Age:  30,
		City: "上海",
	}
	fmt.Printf("p3: %+v\n", p3)

	// 方式4：部分初始化（其他字段为零值）
	p4 := PersonInfo{Name: "王五"}
	fmt.Printf("p4: %+v\n", p4)

	// 访问字段
	fmt.Printf("\np2 的姓名: %s\n", p2.Name)
	fmt.Printf("p2 的年龄: %d\n", p2.Age)

	// 修改字段
	p2.Age = 26
	fmt.Printf("修改后 p2 的年龄: %d\n", p2.Age)

	fmt.Println("\n=== 结构体指针 ===")

	// 使用 new 创建结构体指针
	p5 := new(PersonInfo)
	fmt.Printf("p5: %+v (类型: %T)\n", p5, p5)

	// 通过指针访问字段（Go 自动解引用）
	p5.Name = "赵六"  // 等同于 (*p5).Name = "赵六"
	p5.Age = 28
	fmt.Printf("p5: %+v\n", p5)

	// 获取结构体的指针
	p6 := &PersonInfo{Name: "钱七", Age: 35, City: "深圳"}
	fmt.Printf("p6: %+v\n", p6)

	fmt.Println("\n=== 嵌套结构体 ===")

	pupil := Pupil{
		Name: "小明",
		Age:  18,
		Location: Location{
			Street:  "中关村大街1号",
			City:    "北京",
			ZipCode: "100000",
		},
		Scores: []int{85, 90, 78, 92},
	}

	fmt.Printf("学生信息: %+v\n", pupil)
	fmt.Printf("学生地址: %s, %s\n", pupil.Location.Street, pupil.Location.City)
	fmt.Printf("学生成绩: %v\n", pupil.Scores)

	fmt.Println("\n=== 匿名字段（嵌入） ===")

	// 匿名字段可以直接访问嵌入类型的字段
	member := Member{
		Name: "用户A",
		Contact: Contact{
			Email: "user@example.com",
			Phone: "1234567890",
		},
	}

	fmt.Printf("会员: %+v\n", member)
	// 可以直接访问嵌入字段
	fmt.Printf("邮箱: %s\n", member.Email)  // 等同于 member.Contact.Email
	fmt.Printf("电话: %s\n", member.Phone)  // 等同于 member.Contact.Phone

	fmt.Println("\n=== 结构体比较 ===")

	// 如果结构体的所有字段都是可比较的，则结构体可比较
	a1 := PersonInfo{Name: "测试", Age: 20, City: "北京"}
	a2 := PersonInfo{Name: "测试", Age: 20, City: "北京"}
	a3 := PersonInfo{Name: "测试", Age: 21, City: "北京"}

	fmt.Printf("a1 == a2: %v\n", a1 == a2)
	fmt.Printf("a1 == a3: %v\n", a1 == a3)

	// 注意：包含切片、map 等不可比较字段的结构体无法比较

	fmt.Println("\n=== 匿名结构体 ===")

	// 不定义类型，直接使用匿名结构体
	point := struct {
		X int
		Y int
	}{
		X: 10,
		Y: 20,
	}
	fmt.Printf("点坐标: %+v\n", point)

	// 匿名结构体常用于临时数据或测试
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("配置: %+v\n", config)

	fmt.Println("\n=== 结构体作为函数参数 ===")

	// 值传递（会复制整个结构体）
	showPersonInfo(p3)
	fmt.Printf("调用函数后 p3: %+v (未改变)\n", p3)

	// 指针传递（不复制，可修改原始值）
	modifyPersonInfo(&p3)
	fmt.Printf("调用指针函数后 p3: %+v (已改变)\n", p3)

	fmt.Println("\n=== 结构体数组和切片 ===")

	// 结构体数组
	staffList := [3]Staff{
		{ID: 1, Name: "员工A", Position: "工程师", Salary: 10000, IsActive: true},
		{ID: 2, Name: "员工B", Position: "设计师", Salary: 9000, IsActive: true},
		{ID: 3, Name: "员工C", Position: "经理", Salary: 15000, IsActive: false},
	}

	fmt.Println("员工列表:")
	for _, emp := range staffList {
		fmt.Printf("  ID: %d, 姓名: %s, 职位: %s, 薪资: %.2f\n",
			emp.ID, emp.Name, emp.Position, emp.Salary)
	}

	// 结构体切片
	people := []PersonInfo{
		{Name: "张三", Age: 25, City: "北京"},
		{Name: "李四", Age: 30, City: "上海"},
		{Name: "王五", Age: 28, City: "广州"},
	}

	// 添加新元素
	people = append(people, PersonInfo{Name: "赵六", Age: 32, City: "深圳"})

	fmt.Println("\n人员列表:")
	for i, person := range people {
		fmt.Printf("  %d. %s, %d岁, 来自%s\n", i+1, person.Name, person.Age, person.City)
	}

	fmt.Println("\n=== 空结构体 ===")

	// 空结构体不占用内存空间
	var _ Placeholder  // 使用 _ 忽略未使用的变量
	fmt.Printf("空结构体大小: %d 字节\n", 0)  // 实际为 0

	// 常用于实现集合或信号通道
	set := make(map[string]struct{})
	set["item1"] = struct{}{}
	set["item2"] = struct{}{}
	fmt.Printf("使用空结构体的集合: %v\n", set)
}

// 值传递：接收结构体副本
func showPersonInfo(p PersonInfo) {
	fmt.Printf("函数内: %+v\n", p)
	p.Age = 100  // 修改副本，不影响原始值
}

// 指针传递：接收结构体指针，可修改原始值
func modifyPersonInfo(p *PersonInfo) {
	p.Age = 40  // 修改原始值
	p.City = "杭州"
}
