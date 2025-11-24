package main

import (
	"fmt"
	"math"
)

/*
11 - 方法

学习目标：
1. 学习如何为类型定义方法
2. 理解值接收者和指针接收者的区别
3. 学习方法的链式调用
4. 了解方法集

运行方式：
go run 11_methods.go
*/

// 定义一个矩形结构体
type Rect struct {
	Width  float64
	Height float64
}

// 值接收者方法：计算面积
// 方法不会修改原始结构体
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

// 值接收者方法：计算周长
func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 指针接收者方法：缩放
// 方法会修改原始结构体
func (r *Rect) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// 指针接收者方法：设置尺寸
func (r *Rect) SetDimensions(width, height float64) {
	r.Width = width
	r.Height = height
}

// 圆形结构体
type Circ struct {
	Radius float64
}

// 为 Circ 定义方法
func (c Circ) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circ) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c *Circ) SetRadius(radius float64) {
	c.Radius = radius
}

// 为自定义类型定义方法
type Counter int

func (c *Counter) Increment() {
	*c++
}

func (c *Counter) Decrement() {
	*c--
}

func (c Counter) Value() int {
	return int(c)
}

func (c *Counter) Reset() {
	*c = 0
}

// 字符串切片的自定义类型
type StringList []string

func (s StringList) Print() {
	fmt.Println("字符串列表:")
	for i, str := range s {
		fmt.Printf("  %d: %s\n", i, str)
	}
}

func (s *StringList) Append(str string) {
	*s = append(*s, str)
}

func (s StringList) Length() int {
	return len(s)
}

// 支持链式调用的结构体
type TextBuilder struct {
	data string
}

func (b *TextBuilder) Append(s string) *TextBuilder {
	b.data += s
	return b
}

func (b *TextBuilder) AppendLine(s string) *TextBuilder {
	b.data += s + "\n"
	return b
}

func (b *TextBuilder) Build() string {
	return b.data
}

// 人员结构体，演示方法的组合
type Individual struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Individual) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Individual) HaveBirthday() {
	p.Age++
}

func (p Individual) IsAdult() bool {
	return p.Age >= 18
}

func (p Individual) Greet() {
	fmt.Printf("你好，我是 %s，今年 %d 岁\n", p.FullName(), p.Age)
}

func main() {
	fmt.Println("=== 基本方法调用 ===")

	// 创建矩形
	rect := Rect{Width: 10, Height: 5}
	fmt.Printf("矩形: 宽=%.2f, 高=%.2f\n", rect.Width, rect.Height)
	fmt.Printf("面积: %.2f\n", rect.Area())
	fmt.Printf("周长: %.2f\n", rect.Perimeter())

	fmt.Println("\n=== 值接收者 vs 指针接收者 ===")

	// 值接收者：不会修改原始值
	rect2 := Rect{Width: 10, Height: 5}
	fmt.Printf("缩放前: %+v\n", rect2)

	// 调用指针接收者方法（Go 会自动取地址）
	rect2.Scale(2.0)
	fmt.Printf("缩放后: %+v\n", rect2)

	// 显式使用指针
	rectPtr := &Rect{Width: 20, Height: 10}
	rectPtr.Scale(0.5)
	fmt.Printf("通过指针缩放: %+v\n", rectPtr)

	// Go 的语法糖：值也可以调用指针接收者方法
	rect3 := Rect{Width: 5, Height: 5}
	rect3.SetDimensions(15, 8)  // Go 自动转换为 (&rect3).SetDimensions(15, 8)
	fmt.Printf("设置尺寸后: %+v\n", rect3)

	fmt.Println("\n=== 为不同类型定义方法 ===")

	// 圆形
	circle := Circ{Radius: 5}
	fmt.Printf("圆形半径: %.2f\n", circle.Radius)
	fmt.Printf("圆形面积: %.2f\n", circle.Area())
	fmt.Printf("圆形周长: %.2f\n", circle.Perimeter())

	circle.SetRadius(10)
	fmt.Printf("新半径: %.2f, 新面积: %.2f\n", circle.Radius, circle.Area())

	// 计数器
	var counter Counter = 0
	fmt.Printf("初始计数: %d\n", counter.Value())

	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Printf("增加 3 次: %d\n", counter.Value())

	counter.Decrement()
	fmt.Printf("减少 1 次: %d\n", counter.Value())

	counter.Reset()
	fmt.Printf("重置后: %d\n", counter.Value())

	fmt.Println("\n=== 为切片类型定义方法 ===")

	fruits := StringList{"苹果", "香蕉", "橙子"}
	fruits.Print()

	fmt.Printf("长度: %d\n", fruits.Length())

	fruits.Append("葡萄")
	fruits.Append("西瓜")
	fmt.Println("\n添加后:")
	fruits.Print()

	fmt.Println("\n=== 链式调用 ===")

	// 方法返回自身指针，支持链式调用
	builder := &TextBuilder{}
	result := builder.
		Append("Hello").
		Append(" ").
		Append("World").
		AppendLine("!").
		AppendLine("这是第二行").
		Build()

	fmt.Printf("构建结果:\n%s\n", result)

	fmt.Println("=== 方法的组合使用 ===")

	person := Individual{
		FirstName: "张",
		LastName:  "三",
		Age:       17,
	}

	person.Greet()
	fmt.Printf("是否成年: %v\n", person.IsAdult())

	person.HaveBirthday()
	fmt.Printf("过了生日，年龄: %d\n", person.Age)
	fmt.Printf("现在是否成年: %v\n", person.IsAdult())

	fmt.Println("\n=== 值接收者和指针接收者的选择 ===")

	fmt.Println(`
选择指针接收者的情况：
1. 方法需要修改接收者
2. 接收者是大型结构体（避免复制）
3. 需要保持一致性（如果某些方法用指针接收者，其他方法也应该用）

选择值接收者的情况：
1. 方法不需要修改接收者
2. 接收者是小型结构体或基本类型
3. 接收者是 map、slice、channel 等（它们本身就是引用类型）
`)

	fmt.Println("\n=== 方法集 ===")

	// 类型 T 的方法集包含所有值接收者方法
	// 类型 *T 的方法集包含所有值接收者和指针接收者方法

	var r Rect = Rect{Width: 10, Height: 5}
	var rPtr *Rect = &Rect{Width: 10, Height: 5}

	// 值类型可以调用值接收者和指针接收者方法（Go 会自动转换）
	fmt.Printf("值类型调用: 面积=%.2f\n", r.Area())
	r.Scale(2)  // Go 自动转换为 (&r).Scale(2)
	fmt.Printf("值类型调用指针方法: %+v\n", r)

	// 指针类型可以调用所有方法
	fmt.Printf("指针类型调用: 面积=%.2f\n", rPtr.Area())
	rPtr.Scale(0.5)
	fmt.Printf("指针类型调用指针方法: %+v\n", rPtr)
}
