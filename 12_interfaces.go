package main

import (
	"fmt"
	"math"
)

/*
12 - 接口

学习目标：
1. 学习接口的定义和实现
2. 理解接口的隐式实现
3. 学习空接口和类型断言
4. 掌握接口的组合

运行方式：
go run 12_interfaces.go
*/

// 定义一个形状接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 矩形结构体
type RectShape struct {
	Width  float64
	Height float64
}

// RectShape 实现 Shape 接口
func (r RectShape) Area() float64 {
	return r.Width * r.Height
}

func (r RectShape) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 圆形结构体
type CircShape struct {
	Radius float64
}

// CircShape 实现 Shape 接口
func (c CircShape) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c CircShape) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 三角形结构体
type TriShape struct {
	A, B, C float64
}

// TriShape 实现 Shape 接口
func (t TriShape) Area() float64 {
	// 使用海伦公式
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t TriShape) Perimeter() float64 {
	return t.A + t.B + t.C
}

// 打印形状信息的函数（接受 Shape 接口）
func PrintShapeInfo(s Shape) {
	fmt.Printf("类型: %T\n", s)
	fmt.Printf("面积: %.2f\n", s.Area())
	fmt.Printf("周长: %.2f\n", s.Perimeter())
}

// 定义其他接口

// 命名接口
type Named interface {
	Name() string
}

// 可描述接口
type Describable interface {
	Description() string
}

// 组合接口：同时具有 Named 和 Describable
type Entity interface {
	Named
	Describable
}

// 商品结构体
type ShopItem struct {
	ID          int
	ProductName string
	Price       float64
}

func (p ShopItem) Name() string {
	return p.ProductName
}

func (p ShopItem) Description() string {
	return fmt.Sprintf("产品ID: %d, 价格: ¥%.2f", p.ID, p.Price)
}

// 空接口示例
func printAnything(v interface{}) {
	fmt.Printf("值: %v, 类型: %T\n", v, v)
}

// 类型断言示例
func checkType(v interface{}) {
	// 类型断言：检查接口值的具体类型
	if str, ok := v.(string); ok {
		fmt.Printf("这是一个字符串: %s\n", str)
	} else if num, ok := v.(int); ok {
		fmt.Printf("这是一个整数: %d\n", num)
	} else if shape, ok := v.(Shape); ok {
		fmt.Printf("这是一个形状，面积: %.2f\n", shape.Area())
	} else {
		fmt.Printf("未知类型: %T\n", v)
	}
}

// 使用 type switch
func describeType(v interface{}) {
	switch val := v.(type) {
	case string:
		fmt.Printf("字符串，长度: %d\n", len(val))
	case int:
		fmt.Printf("整数，值: %d\n", val)
	case float64:
		fmt.Printf("浮点数，值: %.2f\n", val)
	case Shape:
		fmt.Printf("形状，面积: %.2f\n", val.Area())
	case []int:
		fmt.Printf("整数切片，长度: %d\n", len(val))
	case nil:
		fmt.Println("nil 值")
	default:
		fmt.Printf("其他类型: %T\n", val)
	}
}

// 定义读取器接口
type DataReader interface {
	Read() string
}

// 定义写入器接口
type DataWriter interface {
	Write(string)
}

// 定义读写器接口（组合）
type DataReadWriter interface {
	DataReader
	DataWriter
}

// 文件结构体
type DataFile struct {
	Content string
}

func (f *DataFile) Read() string {
	return f.Content
}

func (f *DataFile) Write(data string) {
	f.Content = data
}

// 接口的多态性
type Speaker interface {
	Speak() string
}

type DogPet struct {
	Name string
}

func (d DogPet) Speak() string {
	return "汪汪汪"
}

type CatPet struct {
	Name string
}

func (c CatPet) Speak() string {
	return "喵喵喵"
}

type CowPet struct {
	Name string
}

func (c CowPet) Speak() string {
	return "哞哞哞"
}

func main() {
	fmt.Println("=== 基本接口 ===")

	// 创建不同的形状
	rect := RectShape{Width: 10, Height: 5}
	circle := CircShape{Radius: 7}
	triangle := TriShape{A: 3, B: 4, C: 5}

	// 使用接口变量
	var shape Shape

	shape = rect
	PrintShapeInfo(shape)
	fmt.Println()

	shape = circle
	PrintShapeInfo(shape)
	fmt.Println()

	shape = triangle
	PrintShapeInfo(shape)
	fmt.Println()

	// 形状切片（接口类型）
	shapes := []Shape{
		RectShape{Width: 8, Height: 6},
		CircShape{Radius: 5},
		TriShape{A: 5, B: 6, C: 7},
	}

	fmt.Println("所有形状的总面积:")
	totalArea := 0.0
	for i, s := range shapes {
		area := s.Area()
		fmt.Printf("  形状 %d (%T): 面积 = %.2f\n", i+1, s, area)
		totalArea += area
	}
	fmt.Printf("总面积: %.2f\n\n", totalArea)

	fmt.Println("=== 接口组合 ===")

	product := ShopItem{
		ID:          1001,
		ProductName: "笔记本电脑",
		Price:       5999.99,
	}

	// ShopItem 实现了 Entity 接口（通过实现 Named 和 Describable）
	var entity Entity = product
	fmt.Printf("名称: %s\n", entity.Name())
	fmt.Printf("描述: %s\n\n", entity.Description())

	fmt.Println("=== 空接口 ===")

	// interface{} 或 any 可以接受任何类型的值
	printAnything(42)
	printAnything("Hello")
	printAnything(3.14)
	printAnything([]int{1, 2, 3})
	printAnything(rect)
	fmt.Println()

	fmt.Println("=== 类型断言 ===")

	checkType("Hello, Go!")
	checkType(100)
	checkType(CircShape{Radius: 3})
	checkType(3.14)
	fmt.Println()

	fmt.Println("=== Type Switch ===")

	describeType("测试字符串")
	describeType(42)
	describeType(3.14159)
	describeType(RectShape{Width: 5, Height: 10})
	describeType([]int{1, 2, 3, 4, 5})
	describeType(nil)
	fmt.Println()

	fmt.Println("=== 接口组合：DataReadWriter ===")

	file := &DataFile{}

	// DataFile 实现了 DataReadWriter 接口
	var rw DataReadWriter = file

	rw.Write("Hello, Go Interface!")
	fmt.Printf("读取内容: %s\n\n", rw.Read())

	fmt.Println("=== 多态性 ===")

	animals := []Speaker{
		DogPet{Name: "旺财"},
		CatPet{Name: "咪咪"},
		CowPet{Name: "哞哞"},
	}

	fmt.Println("动物们说话:")
	for _, animal := range animals {
		fmt.Printf("  %T: %s\n", animal, animal.Speak())
	}
	fmt.Println()

	fmt.Println("=== 接口值 ===")

	// 接口值包含两部分：类型和值
	var s Shape
	fmt.Printf("空接口: %v, 类型: %T, 是否为 nil: %v\n", s, s, s == nil)

	s = RectShape{Width: 5, Height: 3}
	fmt.Printf("赋值后: %v, 类型: %T, 是否为 nil: %v\n", s, s, s == nil)

	// 接口的零值是 nil
	var r DataReader
	fmt.Printf("DataReader 接口: %v, 是否为 nil: %v\n\n", r, r == nil)

	fmt.Println("=== 常用接口模式 ===")

	fmt.Println(`
Go 标准库中的常用接口：

1. io.Reader - 读取数据
   type Reader interface {
       Read(p []byte) (n int, err error)
   }

2. io.Writer - 写入数据
   type Writer interface {
       Write(p []byte) (n int, err error)
   }

3. fmt.Stringer - 自定义字符串表示
   type Stringer interface {
       String() string
   }

4. error - 错误处理
   type error interface {
       Error() string
   }

5. sort.Interface - 排序
   type Interface interface {
       Len() int
       Less(i, j int) bool
       Swap(i, j int)
   }
	`)

	fmt.Println("接口设计原则:")
	fmt.Println("1. 接口应该小而精（单一职责）")
	fmt.Println("2. 接受接口，返回具体类型")
	fmt.Println("3. 在使用处定义接口，而不是实现处")
	fmt.Println("4. 接口越大，抽象越弱")
}
