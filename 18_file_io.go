package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/*
18 - 文件操作

学习目标：
1. 学习文件的读取和写入
2. 学习文件的创建和删除
3. 学习目录操作
4. 学习 bufio 缓冲 IO

运行方式：
go run 18_file_io.go
*/

func main() {
	fmt.Println("=== 创建和写入文件 ===")

	// 创建文件
	filename := "test_file.txt"

	// 方式1：使用 os.Create
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}

	// 写入内容
	content := "Hello, Go!\n这是第二行\n第三行内容\n"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("写入失败: %v\n", err)
		file.Close()
		return
	}

	// 关闭文件
	file.Close()
	fmt.Printf("文件 %s 创建成功\n\n", filename)

	fmt.Println("=== 读取文件 ===")

	// 方式1：使用 os.ReadFile 一次性读取（小文件）
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}
	fmt.Println("文件内容（os.ReadFile）:")
	fmt.Println(string(data))

	// 方式2：使用 os.Open 和 Read（大文件）
	fmt.Println("使用 os.Open 分块读取:")
	file2, err := os.Open(filename)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file2.Close()

	buffer := make([]byte, 10)  // 每次读取 10 字节
	for {
		n, err := file2.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("读取错误: %v\n", err)
			break
		}
		fmt.Printf("读取了 %d 字节: %s\n", n, string(buffer[:n]))
	}

	fmt.Println("\n=== 使用 bufio 按行读取 ===")

	file3, _ := os.Open(filename)
	defer file3.Close()

	scanner := bufio.NewScanner(file3)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("第 %d 行: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("扫描错误: %v\n", err)
	}

	fmt.Println("\n=== 追加写入 ===")

	// 以追加模式打开文件
	file4, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file4.Close()

	_, err = file4.WriteString("追加的新内容\n")
	if err != nil {
		fmt.Printf("追加失败: %v\n", err)
		return
	}
	fmt.Println("追加成功")

	// 验证追加结果
	data, _ = os.ReadFile(filename)
	fmt.Println("追加后的内容:")
	fmt.Println(string(data))

	fmt.Println("=== 使用 bufio.Writer ===")

	filename2 := "buffered_file.txt"
	file5, _ := os.Create(filename2)
	defer file5.Close()

	writer := bufio.NewWriter(file5)

	// 写入缓冲区
	writer.WriteString("使用 bufio.Writer\n")
	writer.WriteString("缓冲写入更高效\n")
	writer.WriteString("适合大量小写入操作\n")

	// 刷新缓冲区到文件
	writer.Flush()
	fmt.Printf("使用 bufio.Writer 创建了 %s\n\n", filename2)

	fmt.Println("=== 文件信息 ===")

	info, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("获取文件信息失败: %v\n", err)
		return
	}

	fmt.Printf("文件名: %s\n", info.Name())
	fmt.Printf("大小: %d 字节\n", info.Size())
	fmt.Printf("权限: %s\n", info.Mode())
	fmt.Printf("修改时间: %s\n", info.ModTime())
	fmt.Printf("是否是目录: %v\n\n", info.IsDir())

	fmt.Println("=== 检查文件是否存在 ===")

	if exists := fileExists(filename); exists {
		fmt.Printf("%s 存在\n", filename)
	}

	if exists := fileExists("不存在的文件.txt"); !exists {
		fmt.Println("不存在的文件.txt 不存在")
	}

	fmt.Println("\n=== 目录操作 ===")

	// 创建目录
	dirName := "test_dir"
	err = os.Mkdir(dirName, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("创建目录失败: %v\n", err)
	} else {
		fmt.Printf("目录 %s 创建成功\n", dirName)
	}

	// 创建多级目录
	nestedDir := "parent/child/grandchild"
	err = os.MkdirAll(nestedDir, 0755)
	if err != nil {
		fmt.Printf("创建多级目录失败: %v\n", err)
	} else {
		fmt.Printf("多级目录 %s 创建成功\n", nestedDir)
	}

	// 读取目录内容
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("读取目录失败: %v\n", err)
	} else {
		fmt.Println("\n当前目录内容:")
		for _, entry := range entries {
			if entry.IsDir() {
				fmt.Printf("  [目录] %s\n", entry.Name())
			} else {
				fmt.Printf("  [文件] %s\n", entry.Name())
			}
		}
	}

	fmt.Println("\n=== 路径操作 ===")

	path := "/path/to/file.txt"
	fmt.Printf("路径: %s\n", path)
	fmt.Printf("目录: %s\n", filepath.Dir(path))
	fmt.Printf("文件名: %s\n", filepath.Base(path))
	fmt.Printf("扩展名: %s\n", filepath.Ext(path))

	// 路径拼接
	newPath := filepath.Join("dir1", "dir2", "file.go")
	fmt.Printf("拼接路径: %s\n", newPath)

	// 获取绝对路径
	absPath, _ := filepath.Abs(".")
	fmt.Printf("当前绝对路径: %s\n", absPath)

	fmt.Println("\n=== 复制文件 ===")

	err = copyFile(filename, "copied_file.txt")
	if err != nil {
		fmt.Printf("复制失败: %v\n", err)
	} else {
		fmt.Println("文件复制成功")
	}

	fmt.Println("\n=== 遍历目录 ===")

	fmt.Println("遍历当前目录下的所有文件:")
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Printf("  %s (大小: %d 字节)\n", path, info.Size())
		}
		return nil
	})

	fmt.Println("\n=== 清理测试文件 ===")

	// 删除文件
	os.Remove(filename)
	os.Remove(filename2)
	os.Remove("copied_file.txt")

	// 删除目录
	os.Remove(dirName)
	os.RemoveAll("parent")

	fmt.Println("测试文件和目录已清理")

	fmt.Println("\n=== 文件操作最佳实践 ===")

	fmt.Println(`
文件操作最佳实践：

1. 总是处理错误
   file, err := os.Open(filename)
   if err != nil {
       return err
   }

2. 使用 defer 关闭文件
   file, err := os.Open(filename)
   if err != nil { return err }
   defer file.Close()

3. 小文件用 os.ReadFile
   data, err := os.ReadFile("small.txt")

4. 大文件用流式读取
   reader := bufio.NewReader(file)
   for {
       line, err := reader.ReadString('\n')
       ...
   }

5. 使用 bufio 提高效率
   - bufio.Reader 缓冲读取
   - bufio.Writer 缓冲写入
   - 记得 Flush()

6. 路径操作使用 filepath 包
   - 跨平台兼容
   - filepath.Join() 拼接路径

7. 注意文件权限
   - 0644: 所有者读写，其他只读
   - 0755: 目录的常用权限
	`)
}

// 检查文件是否存在
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// 复制文件
func copyFile(src, dst string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 复制内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// 同步到磁盘
	return dstFile.Sync()
}
