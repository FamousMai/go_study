package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

/*
*
 1. 最开始时会有 nil检查， 确保 single Instance单例实例在最开始时为空。
    这是为了防止在每次调用 getInstance 方法时都去执行消耗巨大的锁定操作。
    如果检查不通过， 则就意味着 single Instance字段已被填充。
 2. single Instance结构体将在锁定期间创建。
 3. 在获取到锁后还会有另一个 nil检查。 这是为了确保即便是有多个协程绕过了第一次检查， 也只能有一个可以创建单例实例。
    否则， 所有协程都会创建自己的单例结构体实例。
*/
func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}

func main() {
	var waitGroup sync.WaitGroup

	for i := 0; i < 10; i++ {

		waitGroup.Add(1)
		go func() {
			getInstance()
			defer waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
