package main

import (
	"fmt"
)

/*
20 - 泛型（Go 1.18+）

学习目标：
1. 学习泛型函数的定义和使用
2. 学习类型参数和约束
3. 学习泛型类型
4. 了解类型推断

运行方式：
go run 20_generics.go
*/

func main() {
	fmt.Println("=== 泛型函数基础 ===")

	// 没有泛型时，需要为每种类型写重复的函数
	fmt.Println("传统方式：")
	fmt.Printf("int 最大值: %d\n", MaxInt(10, 20))
	fmt.Printf("float64 最大值: %f\n", MaxFloat64(3.14, 2.71))

	// 使用泛型函数
	fmt.Println("\n泛型方式：")
	fmt.Printf("int 最大值: %d\n", Max(10, 20))
	fmt.Printf("float64 最大值: %f\n", Max(3.14, 2.71))
	fmt.Printf("string 最大值: %s\n", Max("apple", "banana"))

	fmt.Println("\n=== 类型推断 ===")

	// Go 可以推断类型参数
	result1 := Max(100, 200)      // 推断为 int
	result2 := Max(1.5, 2.5)      // 推断为 float64
	result3 := Max("hello", "world")  // 推断为 string

	fmt.Printf("推断类型: %T = %v\n", result1, result1)
	fmt.Printf("推断类型: %T = %v\n", result2, result2)
	fmt.Printf("推断类型: %T = %v\n", result3, result3)

	// 也可以显式指定类型
	result4 := Max[int](50, 30)
	fmt.Printf("显式类型: %T = %v\n", result4, result4)

	fmt.Println("\n=== 多类型参数 ===")

	pair := MakePair("name", 42)
	fmt.Printf("Pair: %+v\n", pair)
	fmt.Printf("First: %v, Second: %v\n", pair.First, pair.Second)

	fmt.Println("\n=== 泛型切片函数 ===")

	// 通用的 Map 函数
	numbers := []int{1, 2, 3, 4, 5}
	doubled := Map(numbers, func(n int) int {
		return n * 2
	})
	fmt.Printf("原始: %v\n", numbers)
	fmt.Printf("翻倍: %v\n", doubled)

	// 类型转换
	strings := Map(numbers, func(n int) string {
		return fmt.Sprintf("#%d", n)
	})
	fmt.Printf("转换: %v\n", strings)

	// 通用的 Filter 函数
	evens := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("偶数: %v\n", evens)

	// 通用的 Reduce 函数
	sum := Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Printf("求和: %d\n", sum)

	fmt.Println("\n=== 泛型类型 ===")

	// 泛型栈
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Printf("栈大小: %d\n", intStack.Size())
	fmt.Printf("弹出: %d\n", intStack.Pop())
	fmt.Printf("弹出: %d\n", intStack.Pop())

	stringStack := &Stack[string]{}
	stringStack.Push("a")
	stringStack.Push("b")
	fmt.Printf("字符串栈顶: %s\n", stringStack.Pop())

	fmt.Println("\n=== 类型约束 ===")

	// 使用 any 约束（等价于 interface{}）
	PrintValue(42)
	PrintValue("hello")
	PrintValue(3.14)

	// 使用 comparable 约束
	fmt.Printf("Contains: %v\n", Contains([]int{1, 2, 3, 4, 5}, 3))
	fmt.Printf("Contains: %v\n", Contains([]string{"a", "b", "c"}, "d"))

	// 使用自定义约束
	fmt.Printf("Sum: %d\n", Sum([]int{1, 2, 3, 4, 5}))
	fmt.Printf("Sum: %f\n", Sum([]float64{1.1, 2.2, 3.3}))

	fmt.Println("\n=== 泛型接口约束 ===")

	// 约束必须实现某个接口
	dog := Dog{Name: "旺财"}
	cat := Cat{Name: "咪咪"}

	Speak(dog)
	Speak(cat)

	fmt.Println("\n=== 泛型最佳实践 ===")

	fmt.Println(`
泛型最佳实践：

1. 何时使用泛型
   - 需要处理多种类型的相同逻辑
   - 容器类型（栈、队列、树等）
   - 通用算法（排序、查找等）

2. 何时不使用泛型
   - 逻辑只适用于特定类型
   - 简单的类型转换
   - 已有的接口可以解决问题

3. 类型约束选择
   - any: 任何类型
   - comparable: 支持 == 和 != 的类型
   - 自定义约束: 需要特定方法或操作的类型

4. 性能考虑
   - 泛型代码在编译时实例化
   - 对于基本类型，性能与手写代码相当
   - 过度使用可能增加编译时间

5. 可读性
   - 不要为了使用泛型而使用泛型
   - 保持代码简洁清晰
   - 为类型参数选择有意义的名称
	`)
}

// ========== 传统的非泛型函数 ==========

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// ========== 泛型函数 ==========

// Ordered 约束：支持比较操作的类型
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64 |
	~string
}

// 泛型 Max 函数
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// ========== 多类型参数 ==========

type Pair[T, U any] struct {
	First  T
	Second U
}

func MakePair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

// ========== 泛型切片函数 ==========

// Map 函数
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter 函数
func Filter[T any](slice []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce 函数
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// ========== 泛型类型 ==========

// 泛型栈
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

// ========== 类型约束 ==========

// 使用 any 约束
func PrintValue[T any](v T) {
	fmt.Printf("值: %v (类型: %T)\n", v, v)
}

// 使用 comparable 约束
func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// 自定义数值约束
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64
}

func Sum[T Number](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// ========== 接口约束 ==========

type Animal interface {
	SayHello() string
}

type Dog struct {
	Name string
}

func (d Dog) SayHello() string {
	return fmt.Sprintf("%s: 汪汪汪!", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) SayHello() string {
	return fmt.Sprintf("%s: 喵喵喵!", c.Name)
}

// 约束必须实现 Animal 接口
func Speak[T Animal](a T) {
	fmt.Println(a.SayHello())
}
