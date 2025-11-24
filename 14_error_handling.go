package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
14 - 错误处理

学习目标：
1. 理解 Go 的错误处理机制
2. 学习如何创建和返回错误
3. 学习错误处理的最佳实践
4. 了解自定义错误类型

运行方式：
go run 14_error_handling.go
*/

func main() {
	fmt.Println("=== 基本错误处理 ===")

	// 调用可能返回错误的函数
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	// 除以零的情况
	result2, err2 := safeDivide(10, 0)
	if err2 != nil {
		fmt.Printf("错误: %v\n", err2)
	} else {
		fmt.Printf("结果: %.2f\n", result2)
	}

	fmt.Println("\n=== 创建错误 ===")

	// 方式1：使用 errors.New
	err1 := errors.New("这是一个错误")
	fmt.Printf("err1: %v\n", err1)

	// 方式2：使用 fmt.Errorf（支持格式化）
	username := "admin"
	err3 := fmt.Errorf("用户 %s 不存在", username)
	fmt.Printf("err2: %v\n", err3)

	fmt.Println("\n=== 错误处理模式 ===")

	// 模式1：立即处理错误
	age, err := parseAgeStr("25")
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
		return // 或其他错误处理
	}
	fmt.Printf("年龄: %d\n", age)

	// 模式2：错误传播
	user, err := fetchUserData("user123")
	if err != nil {
		fmt.Printf("获取用户信息失败: %v\n", err)
	} else {
		fmt.Printf("用户: %+v\n", user)
	}

	fmt.Println("\n=== 哨兵错误（Sentinel Errors） ===")

	// 定义预定义的错误值
	var (
		ErrNotFound = errors.New("未找到")
		// ErrInvalid 可用于其他验证场景
	)

	err = ErrNotFound
	if err == ErrNotFound {
		fmt.Println("这是一个 NotFound 错误")
	}

	// 使用标准库的哨兵错误
	_, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Printf("转换错误: %v\n", err)
		// 错误类型断言
		if numErr, ok := err.(*strconv.NumError); ok {
			fmt.Printf("  错误类型: %s\n", numErr.Err)
		}
	}

	fmt.Println("\n=== 自定义错误类型 ===")

	// 使用自定义错误类型
	err = checkAge(-5)
	if err != nil {
		fmt.Printf("验证错误: %v\n", err)
		// 类型断言获取详细信息
		if validationErr, ok := err.(*FieldValidationError); ok {
			fmt.Printf("  字段: %s\n", validationErr.Field)
			fmt.Printf("  原因: %s\n", validationErr.Reason)
		}
	}

	err = checkAge(25)
	if err == nil {
		fmt.Println("年龄验证通过")
	}

	fmt.Println("\n=== 错误包装（Error Wrapping） ===")

	// Go 1.13+ 支持错误包装
	err = handleFile("config.txt")
	if err != nil {
		fmt.Printf("处理文件失败: %v\n", err)
	}

	fmt.Println("\n=== 多返回值错误处理 ===")

	// 返回多个值和错误
	width, height, err := parseDimensions("10x20")
	if err != nil {
		fmt.Printf("解析失败: %v\n", err)
	} else {
		fmt.Printf("宽度: %d, 高度: %d\n", width, height)
	}

	_, _, err = parseDimensions("invalid")
	if err != nil {
		fmt.Printf("解析失败: %v\n", err)
	}

	fmt.Println("\n=== 优雅的错误处理 ===")

	// 示例：读取配置
	config, err := loadAppConfig("app.conf")
	if err != nil {
		fmt.Printf("加载配置失败: %v，使用默认配置\n", err)
		config = getDefaultAppConfig()
	}
	fmt.Printf("配置: %+v\n", config)

	fmt.Println("\n=== defer 与错误处理 ===")

	err = executeWithCleanup()
	if err != nil {
		fmt.Printf("操作失败: %v\n", err)
	}

	fmt.Println("\n=== 错误处理最佳实践 ===")

	fmt.Println(`
错误处理最佳实践：

1. 总是检查错误
   if err != nil {
       // 处理错误
   }

2. 错误信息应该清晰、具体
   ❌ errors.New("error")
   ✅ fmt.Errorf("failed to open file %s: %w", filename, err)

3. 不要忽略错误
   ❌ result, _ := someFunc()
   ✅ result, err := someFunc()
      if err != nil { ... }

4. 及早返回错误
   if err != nil {
       return fmt.Errorf("operation failed: %w", err)
   }

5. 为公共 API 提供有意义的错误
   使用自定义错误类型或哨兵错误

6. 在适当的层级处理错误
   - 底层：创建和返回错误
   - 中层：包装和传递错误
   - 顶层：处理和记录错误

7. 使用 %w 包装错误（Go 1.13+）
   return fmt.Errorf("context: %w", originalErr)

8. 不要使用 panic 来处理正常的错误
   panic 应该只用于不可恢复的错误
	`)
}

// 基本错误返回
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// 解析年龄
func parseAgeStr(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("无法解析年龄 '%s': %w", s, err)
	}
	if age < 0 || age > 150 {
		return 0, fmt.Errorf("年龄 %d 超出有效范围", age)
	}
	return age, nil
}

// 用户信息结构体
type UserData struct {
	ID   string
	Name string
	Age  int
}

// 获取用户信息（错误传播示例）
func fetchUserData(userID string) (*UserData, error) {
	// 模拟从数据库获取
	user, err := queryUserFromDB(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info for %s: %w", userID, err)
	}
	return user, nil
}

func queryUserFromDB(userID string) (*UserData, error) {
	// 模拟数据库查询
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	// 模拟找不到用户
	return nil, errors.New("user not found in database")
}

// 自定义错误类型
type FieldValidationError struct {
	Field  string
	Reason string
}

func (e *FieldValidationError) Error() string {
	return fmt.Sprintf("验证失败: 字段 %s, 原因: %s", e.Field, e.Reason)
}

// 使用自定义错误类型
func checkAge(age int) error {
	if age < 0 {
		return &FieldValidationError{
			Field:  "age",
			Reason: "年龄不能为负数",
		}
	}
	if age > 150 {
		return &FieldValidationError{
			Field:  "age",
			Reason: "年龄不能超过150",
		}
	}
	return nil
}

// 错误包装示例
func handleFile(filename string) error {
	err := loadFile(filename)
	if err != nil {
		return fmt.Errorf("处理文件 %s 失败: %w", filename, err)
	}
	return nil
}

func loadFile(filename string) error {
	// 模拟文件读取错误
	return errors.New("文件不存在")
}

// 解析尺寸
func parseDimensions(s string) (width, height int, err error) {
	// 简化的解析逻辑
	if s != "10x20" {
		return 0, 0, fmt.Errorf("无效的尺寸格式: %s", s)
	}
	return 10, 20, nil
}

// 配置结构体
type AppConfig struct {
	Host string
	Port int
}

func loadAppConfig(filename string) (*AppConfig, error) {
	// 模拟加载失败
	return nil, fmt.Errorf("无法读取配置文件 %s", filename)
}

func getDefaultAppConfig() *AppConfig {
	return &AppConfig{
		Host: "localhost",
		Port: 8080,
	}
}

// defer 与错误处理
func executeWithCleanup() error {
	// 模拟打开资源
	fmt.Println("  打开资源...")

	// 使用 defer 确保资源被清理
	defer func() {
		fmt.Println("  清理资源...")
	}()

	// 模拟操作
	fmt.Println("  执行操作...")

	// 模拟错误
	return errors.New("操作过程中发生错误")

	// defer 的清理代码仍然会执行
}
