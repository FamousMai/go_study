package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

var once sync.Once

type single struct{}

var singleInstance *single
var singleInstance1 *single

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

/*
*
1. init函数
我们可以在 init函数中创建单例实例。
这仅适用于实例的早期初始化工作已经确定时。
init函数仅会在包中的每个文件里调用一次， 所以我们可以确定其只会创建一个实例。

2. sync.Once
sync.Once仅会执行一次操作。 可查看下面的代码：
*/
func getInstanceUseSyncOnce(i int) *single {
	if singleInstance1 == nil {
		fmt.Println(i)
		once.Do(
			func() {
				fmt.Println("Creating1 single instance now.")
				singleInstance1 = &single{}
			})
	} else {
		fmt.Println("Single instance1 already created.")
	}
	return singleInstance1
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

	var waitGroup2 sync.WaitGroup
	for i := 0; i < 10; i++ {

		waitGroup2.Add(1)
		go func() {
			getInstanceUseSyncOnce(i)
			defer waitGroup2.Done()

		}()
	}

	waitGroup2.Wait()
}
