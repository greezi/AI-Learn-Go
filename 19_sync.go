package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
19 - 并发同步

学习目标：
1. 学习 sync.WaitGroup 等待一组 goroutine
2. 学习 sync.Mutex 互斥锁
3. 学习 sync.RWMutex 读写锁
4. 学习 sync.Once 单次执行
5. 学习原子操作

运行方式：
go run 19_sync.go
*/

func main() {
	fmt.Println("=== sync.WaitGroup ===")

	var wg sync.WaitGroup

	// 启动 5 个 goroutine
	for i := 1; i <= 5; i++ {
		wg.Add(1)  // 计数器加 1
		go func(id int) {
			defer wg.Done()  // 完成时计数器减 1
			fmt.Printf("Goroutine %d 开始\n", id)
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Goroutine %d 结束\n", id)
		}(i)
	}

	fmt.Println("等待所有 goroutine 完成...")
	wg.Wait()  // 阻塞直到计数器为 0
	fmt.Println("所有 goroutine 完成\n")

	fmt.Println("=== sync.Mutex（互斥锁） ===")

	// 没有锁的情况（可能出现数据竞争）
	var counter int
	var wg2 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			counter++  // 不安全的操作
		}()
	}
	wg2.Wait()
	fmt.Printf("没有锁的计数器（可能不准确）: %d\n", counter)

	// 使用互斥锁
	counter = 0
	var mutex sync.Mutex
	var wg3 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			mutex.Lock()    // 加锁
			counter++
			mutex.Unlock()  // 解锁
		}()
	}
	wg3.Wait()
	fmt.Printf("使用互斥锁的计数器: %d\n\n", counter)

	fmt.Println("=== sync.RWMutex（读写锁） ===")

	// 读写锁允许多个读操作同时进行，但写操作独占
	cache := &SafeCache{data: make(map[string]string)}

	var wg4 sync.WaitGroup

	// 启动多个写入者
	for i := 1; i <= 3; i++ {
		wg4.Add(1)
		go func(id int) {
			defer wg4.Done()
			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)
			cache.Set(key, value)
			fmt.Printf("写入: %s = %s\n", key, value)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	// 启动多个读取者
	for i := 1; i <= 5; i++ {
		wg4.Add(1)
		go func(id int) {
			defer wg4.Done()
			key := fmt.Sprintf("key%d", (id%3)+1)
			value := cache.Get(key)
			fmt.Printf("读取: %s = %s\n", key, value)
		}(i)
	}

	wg4.Wait()
	fmt.Println()

	fmt.Println("=== sync.Once ===")

	var once sync.Once
	var wg5 sync.WaitGroup

	// 初始化函数
	initFunc := func() {
		fmt.Println("初始化只执行一次")
	}

	// 多个 goroutine 尝试执行
	for i := 1; i <= 5; i++ {
		wg5.Add(1)
		go func(id int) {
			defer wg5.Done()
			fmt.Printf("Goroutine %d 尝试执行初始化\n", id)
			once.Do(initFunc)  // 只有第一次调用会执行
			fmt.Printf("Goroutine %d 继续执行\n", id)
		}(i)
	}

	wg5.Wait()
	fmt.Println()

	fmt.Println("=== sync/atomic（原子操作） ===")

	var atomicCounter int64

	var wg6 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg6.Add(1)
		go func() {
			defer wg6.Done()
			atomic.AddInt64(&atomicCounter, 1)  // 原子加法
		}()
	}
	wg6.Wait()
	fmt.Printf("原子计数器: %d\n", atomic.LoadInt64(&atomicCounter))

	// 其他原子操作
	var value int64 = 100

	// 原子加载
	loaded := atomic.LoadInt64(&value)
	fmt.Printf("原子加载: %d\n", loaded)

	// 原子存储
	atomic.StoreInt64(&value, 200)
	fmt.Printf("原子存储后: %d\n", value)

	// 原子交换
	old := atomic.SwapInt64(&value, 300)
	fmt.Printf("原子交换: 旧值=%d, 新值=%d\n", old, value)

	// 原子比较并交换（CAS）
	swapped := atomic.CompareAndSwapInt64(&value, 300, 400)
	fmt.Printf("CAS 成功: %v, 当前值: %d\n\n", swapped, value)

	fmt.Println("=== sync.Cond（条件变量） ===")

	condDemo()

	fmt.Println("\n=== sync.Map ===")

	// sync.Map 是并发安全的 map
	var sm sync.Map

	// 存储
	sm.Store("key1", "value1")
	sm.Store("key2", "value2")

	// 读取
	if value, ok := sm.Load("key1"); ok {
		fmt.Printf("key1 = %v\n", value)
	}

	// 读取或存储（如果不存在则存储）
	actual, wasLoaded := sm.LoadOrStore("key3", "value3")
	fmt.Printf("key3 = %v, 已存在: %v\n", actual, wasLoaded)

	// 遍历
	fmt.Println("sync.Map 内容:")
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v = %v\n", key, value)
		return true  // 返回 true 继续遍历
	})

	// 删除
	sm.Delete("key1")
	fmt.Println()

	fmt.Println("=== 同步原语选择指南 ===")

	fmt.Println(`
同步原语选择指南：

1. sync.WaitGroup
   - 等待一组 goroutine 完成
   - 不需要传递数据

2. sync.Mutex
   - 保护共享资源的独占访问
   - 临界区代码需要互斥执行

3. sync.RWMutex
   - 读多写少的场景
   - 允许多个并发读取

4. sync.Once
   - 确保代码只执行一次
   - 单例模式、延迟初始化

5. sync/atomic
   - 简单的计数器操作
   - 无需复杂的锁逻辑

6. sync.Map
   - 高并发读写 map
   - 比 map + Mutex 更高效

7. sync.Cond
   - goroutine 之间的信号通知
   - 等待特定条件满足

8. Channel
   - 优先使用 channel 进行 goroutine 通信
   - "不要通过共享内存来通信，而应该通过通信来共享内存"
	`)
}

// 使用 RWMutex 的安全缓存
type SafeCache struct {
	data  map[string]string
	mutex sync.RWMutex
}

func (c *SafeCache) Set(key, value string) {
	c.mutex.Lock()         // 写锁
	defer c.mutex.Unlock()
	c.data[key] = value
}

func (c *SafeCache) Get(key string) string {
	c.mutex.RLock()        // 读锁
	defer c.mutex.RUnlock()
	return c.data[key]
}

// 条件变量示例
func condDemo() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	ready := false
	var wg sync.WaitGroup

	// 等待者
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			mutex.Lock()
			for !ready {  // 使用循环检查条件
				fmt.Printf("等待者 %d 等待中...\n", id)
				cond.Wait()  // 等待信号
			}
			fmt.Printf("等待者 %d 收到信号\n", id)
			mutex.Unlock()
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	// 通知者
	go func() {
		mutex.Lock()
		ready = true
		fmt.Println("条件已满足，广播信号")
		cond.Broadcast()  // 通知所有等待者
		mutex.Unlock()
	}()

	wg.Wait()
}
