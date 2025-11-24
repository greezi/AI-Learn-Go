package main

import (
	"encoding/json"
	"fmt"
)

/*
21 - JSON 处理

学习目标：
1. 学习 JSON 的编码（Marshal）
2. 学习 JSON 的解码（Unmarshal）
3. 学习结构体标签
4. 学习自定义 JSON 处理

运行方式：
go run 21_json.go
*/

// 账户结构体（带 JSON 标签）
type Account struct {
	ID        int      `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email,omitempty"` // omitempty: 空值不输出
	Password  string   `json:"-"`               // -: 忽略该字段
	Age       int      `json:"age"`
	IsActive  bool     `json:"is_active"`
	Tags      []string `json:"tags,omitempty"`
}

// 文章结构体（嵌套）
type BlogPost struct {
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Author  Account `json:"author"`
}

// 自定义 JSON 序列化
type TaskStatus int

const (
	TaskStatusPending TaskStatus = iota
	TaskStatusActive
	TaskStatusInactive
)

func (s TaskStatus) MarshalJSON() ([]byte, error) {
	var str string
	switch s {
	case TaskStatusPending:
		str = "pending"
	case TaskStatusActive:
		str = "active"
	case TaskStatusInactive:
		str = "inactive"
	default:
		str = "unknown"
	}
	return json.Marshal(str)
}

func (s *TaskStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	switch str {
	case "pending":
		*s = TaskStatusPending
	case "active":
		*s = TaskStatusActive
	case "inactive":
		*s = TaskStatusInactive
	default:
		*s = TaskStatusPending
	}
	return nil
}

type TodoTask struct {
	Name   string     `json:"name"`
	Status TaskStatus `json:"status"`
}

func main() {
	fmt.Println("=== JSON 编码（Marshal） ===")

	// 创建账户
	account := Account{
		ID:       1,
		Username: "zhangsan",
		Email:    "zhangsan@example.com",
		Password: "secret123",
		Age:      25,
		IsActive: true,
		Tags:     []string{"go", "developer"},
	}

	// 编码为 JSON
	jsonData, err := json.Marshal(account)
	if err != nil {
		fmt.Printf("编码失败: %v\n", err)
		return
	}
	fmt.Printf("JSON: %s\n", jsonData)

	// 格式化输出（缩进）
	jsonPretty, _ := json.MarshalIndent(account, "", "  ")
	fmt.Printf("格式化 JSON:\n%s\n\n", jsonPretty)

	fmt.Println("=== omitempty 和 - 标签 ===")

	// 空值会被 omitempty 忽略
	account2 := Account{
		ID:       2,
		Username: "lisi",
		IsActive: false,
		// Email 和 Tags 为空，会被忽略
	}
	jsonData2, _ := json.MarshalIndent(account2, "", "  ")
	fmt.Printf("带 omitempty 的 JSON:\n%s\n\n", jsonData2)

	fmt.Println("=== JSON 解码（Unmarshal） ===")

	jsonStr := `{
		"id": 3,
		"username": "wangwu",
		"email": "wangwu@example.com",
		"age": 30,
		"is_active": true,
		"tags": ["backend", "frontend"]
	}`

	var decodedAccount Account
	err = json.Unmarshal([]byte(jsonStr), &decodedAccount)
	if err != nil {
		fmt.Printf("解码失败: %v\n", err)
		return
	}
	fmt.Printf("解码后的账户: %+v\n\n", decodedAccount)

	fmt.Println("=== 嵌套结构体 ===")

	post := BlogPost{
		Title:   "Go 语言入门",
		Content: "这是一篇关于 Go 的文章...",
		Author:  account,
	}

	postJSON, _ := json.MarshalIndent(post, "", "  ")
	fmt.Printf("嵌套结构体 JSON:\n%s\n\n", postJSON)

	fmt.Println("=== 解码到 map ===")

	// 当结构未知时，可以解码到 map
	jsonStr2 := `{"name": "test", "count": 42, "enabled": true}`
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonStr2), &result)

	fmt.Printf("解码到 map: %v\n", result)
	fmt.Printf("name = %v (类型: %T)\n", result["name"], result["name"])
	fmt.Printf("count = %v (类型: %T)\n", result["count"], result["count"])

	fmt.Println("\n=== 解码数组 ===")

	jsonArray := `[
		{"id": 1, "username": "user1"},
		{"id": 2, "username": "user2"},
		{"id": 3, "username": "user3"}
	]`

	var accounts []Account
	json.Unmarshal([]byte(jsonArray), &accounts)

	fmt.Println("账户列表:")
	for _, a := range accounts {
		fmt.Printf("  ID: %d, Username: %s\n", a.ID, a.Username)
	}

	fmt.Println("\n=== 自定义 JSON 序列化 ===")

	task := TodoTask{
		Name:   "完成项目",
		Status: TaskStatusActive,
	}

	taskJSON, _ := json.MarshalIndent(task, "", "  ")
	fmt.Printf("自定义序列化:\n%s\n", taskJSON)

	// 解码
	taskJSON2 := `{"name": "新任务", "status": "pending"}`
	var task2 TodoTask
	json.Unmarshal([]byte(taskJSON2), &task2)
	fmt.Printf("自定义反序列化: %+v (Status: %d)\n\n", task2, task2.Status)

	fmt.Println("=== 处理未知字段 ===")

	// 使用 json.RawMessage 延迟解析
	type ApiResponse struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"` // 延迟解析
	}

	resp := `{"type": "account", "data": {"id": 1, "username": "test"}}`
	var r ApiResponse
	json.Unmarshal([]byte(resp), &r)

	fmt.Printf("类型: %s\n", r.Type)
	fmt.Printf("原始数据: %s\n", r.Data)

	// 根据类型解析 data
	if r.Type == "account" {
		var a Account
		json.Unmarshal(r.Data, &a)
		fmt.Printf("解析的账户: %+v\n", a)
	}

	fmt.Println("\n=== JSON 最佳实践 ===")

	fmt.Println(`
JSON 最佳实践：

1. 总是使用结构体标签
   type Account struct {
       Name string ` + "`json:\"name\"`" + `
   }

2. 使用 omitempty 避免空值
   Email string ` + "`json:\"email,omitempty\"`" + `

3. 使用 - 忽略敏感字段
   Password string ` + "`json:\"-\"`" + `

4. 处理错误
   if err := json.Unmarshal(data, &v); err != nil {
       return err
   }

5. 使用 json.Number 处理数字精度
   var result map[string]json.Number

6. 验证必填字段
   if account.Name == "" {
       return errors.New("name is required")
   }

7. 使用 Decoder/Encoder 处理流
   decoder := json.NewDecoder(reader)
   encoder := json.NewEncoder(writer)
	`)
}
