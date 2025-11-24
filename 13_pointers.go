package main

import "fmt"

/*
13 - 指针

学习目标：
1. 理解指针的概念
2. 学习指针的声明和使用
3. 学习指针与函数
4. 理解指针与结构体

运行方式：
go run 13_pointers.go
*/

func main() {
	fmt.Println("=== 指针基础 ===")

	// 普通变量
	num := 42
	fmt.Printf("num 的值: %d\n", num)
	fmt.Printf("num 的地址: %p\n", &num)  // & 取地址运算符

	// 声明指针变量
	var ptr *int  // ptr 是一个指向 int 的指针
	fmt.Printf("未初始化的指针: %v\n", ptr)  // nil

	// 将 num 的地址赋给指针
	ptr = &num
	fmt.Printf("ptr 的值（num 的地址）: %p\n", ptr)
	fmt.Printf("ptr 指向的值: %d\n", *ptr)  // * 解引用运算符

	// 通过指针修改值
	*ptr = 100
	fmt.Printf("通过指针修改后，num 的值: %d\n", num)

	fmt.Println("\n=== 指针的零值 ===")

	var p *int
	fmt.Printf("指针 p 的零值: %v\n", p)
	fmt.Printf("p 是否为 nil: %v\n", p == nil)

	// 注意：不能对 nil 指针解引用，会 panic
	// fmt.Println(*p)  // 这会导致 panic

	fmt.Println("\n=== new 函数 ===")

	// new 函数分配内存并返回指针
	p2 := new(int)
	fmt.Printf("new(int) 返回的指针: %p\n", p2)
	fmt.Printf("指针指向的值（零值）: %d\n", *p2)

	*p2 = 200
	fmt.Printf("赋值后的值: %d\n", *p2)

	// new 可以用于任何类型
	p3 := new(string)
	*p3 = "Hello"
	fmt.Printf("new(string): %s\n", *p3)

	fmt.Println("\n=== 指针与函数 ===")

	x := 10
	fmt.Printf("调用前 x 的值: %d\n", x)

	// 值传递：不会修改原始值
	incrementValue(x)
	fmt.Printf("值传递后 x 的值: %d (未改变)\n", x)

	// 指针传递：会修改原始值
	incrementPointer(&x)
	fmt.Printf("指针传递后 x 的值: %d (已改变)\n", x)

	fmt.Println("\n=== 指针与结构体 ===")

	// 定义结构体
	type Person struct {
		Name string
		Age  int
	}

	// 创建结构体
	person1 := Person{Name: "张三", Age: 25}
	fmt.Printf("person1: %+v\n", person1)

	// 获取结构体的指针
	personPtr := &person1
	fmt.Printf("personPtr: %p\n", personPtr)

	// 通过指针访问字段（Go 自动解引用）
	fmt.Printf("通过指针访问 Name: %s\n", personPtr.Name)  // 等同于 (*personPtr).Name

	// 通过指针修改字段
	personPtr.Age = 26
	fmt.Printf("修改后 person1.Age: %d\n", person1.Age)

	// 使用 new 创建结构体指针
	person2 := new(Person)
	person2.Name = "李四"
	person2.Age = 30
	fmt.Printf("person2: %+v\n", person2)

	// 值传递结构体
	person3 := PersonForPointer{Name: "王五", Age: 35}
	updatePersonValueLocal(person3)
	fmt.Printf("值传递后 person3: %+v (未改变)\n", person3)

	// 指针传递结构体
	updatePersonPointerLocal(&person3)
	fmt.Printf("指针传递后 person3: %+v (已改变)\n", person3)

	fmt.Println("\n=== 指针数组和数组指针 ===")

	// 指针数组：数组的元素是指针
	a, b, c := 1, 2, 3
	ptrArray := [3]*int{&a, &b, &c}
	fmt.Println("指针数组:")
	for i, ptr := range ptrArray {
		fmt.Printf("  索引 %d: 地址 %p, 值 %d\n", i, ptr, *ptr)
	}

	// 修改指针指向的值
	*ptrArray[0] = 10
	fmt.Printf("修改后 a 的值: %d\n", a)

	// 数组指针：指向数组的指针
	arr := [3]int{10, 20, 30}
	var arrPtr *[3]int = &arr
	fmt.Printf("\n数组指针: %p\n", arrPtr)
	fmt.Printf("数组指针指向的数组: %v\n", *arrPtr)
	fmt.Printf("访问元素: arrPtr[1] = %d\n", arrPtr[1])  // Go 自动解引用

	fmt.Println("\n=== 指针与切片 ===")

	// 切片本身就包含指针（指向底层数组）
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice1: %v\n", slice1)

	// 切片作为函数参数时，可以修改元素（因为切片包含指针）
	modifySlice(slice1)
	fmt.Printf("修改后 slice1: %v (已改变)\n", slice1)

	// 但如果要修改切片本身（如 append），需要传指针
	appendSlice(&slice1)
	fmt.Printf("append 后 slice1: %v\n", slice1)

	fmt.Println("\n=== 指针与 Map ===")

	// map 也是引用类型
	map1 := map[string]int{"a": 1, "b": 2}
	fmt.Printf("map1: %v\n", map1)

	modifyMap(map1)
	fmt.Printf("修改后 map1: %v (已改变)\n", map1)

	fmt.Println("\n=== 多级指针 ===")

	value := 42
	ptr1 := &value      // 指向 int 的指针
	ptr2 := &ptr1       // 指向指针的指针

	fmt.Printf("value: %d\n", value)
	fmt.Printf("ptr1 指向的值: %d\n", *ptr1)
	fmt.Printf("ptr2 指向的指针指向的值: %d\n", **ptr2)

	// 通过二级指针修改值
	**ptr2 = 100
	fmt.Printf("修改后 value: %d\n", value)

	fmt.Println("\n=== 指针的比较 ===")

	num1 := 10
	num2 := 10
	ptrA := &num1
	ptrB := &num2
	ptrC := &num1

	fmt.Printf("ptrA == ptrB: %v (指向不同变量)\n", ptrA == ptrB)
	fmt.Printf("ptrA == ptrC: %v (指向同一变量)\n", ptrA == ptrC)

	fmt.Println("\n=== 指针的实用场景 ===")

	fmt.Println(`
指针的使用场景：

1. 需要修改函数外部的变量
2. 避免复制大型结构体（性能优化）
3. 实现可选参数（使用 nil 指针）
4. 在方法中修改接收者
5. 实现数据结构（链表、树等）

注意事项：

1. 不要返回局部变量的指针给外部使用（Go 会自动处理，但要理解）
2. 避免指针的过度使用，影响代码可读性
3. nil 指针解引用会导致 panic
4. Go 的垃圾回收会自动管理内存，无需手动释放
	`)

	fmt.Println("\n=== 指针性能示例 ===")

	large := LargeStructForPointer{}

	// 值传递：复制整个结构体（开销大）
	processByValueLocal(large)

	// 指针传递：只复制指针（开销小）
	processByPointerLocal(&large)

	fmt.Println("对于大型结构体，使用指针传递性能更好")
}

// 值传递：不会修改原始值
func incrementValue(n int) {
	n++
	fmt.Printf("  函数内 n 的值: %d\n", n)
}

// 指针传递：会修改原始值
func incrementPointer(n *int) {
	*n++
	fmt.Printf("  函数内 *n 的值: %d\n", *n)
}

// 定义结构体用于测试
type PersonForPointer struct {
	Name string
	Age  int
}

// 值传递结构体（本地使用）
func updatePersonValueLocal(p PersonForPointer) {
	p.Age = 100
}

// 指针传递结构体（本地使用）
func updatePersonPointerLocal(p *PersonForPointer) {
	p.Age = 100
}

// 修改切片元素
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
}

// 修改切片本身（append）
func appendSlice(s *[]int) {
	*s = append(*s, 100, 200)
}

// 修改 map
func modifyMap(m map[string]int) {
	m["c"] = 3
}

// 定义大型结构体用于性能测试
type LargeStructForPointer struct {
	Data [1000]int
}

// 值传递大型结构体（本地使用）
func processByValueLocal(l LargeStructForPointer) {
	// 处理数据
}

// 指针传递大型结构体（本地使用）
func processByPointerLocal(l *LargeStructForPointer) {
	// 处理数据
}
