package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构管理队列
type Queue struct {
	maxSize int
	array   [5]int //数组模拟队列
	front   int    //指向队列首
	rear    int    //指向队列尾部
}

func (this *Queue) AddQueue(val int) (err error) {

	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}

	this.rear++
	this.array[this.rear] = val
	return
}

func (this *Queue) ShowQueue() {
	// this.front 不包含对首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("arrar[%d]=%d\t\n", i, this.array[i])
	}
}

func (this *Queue) GetQueue() (val int, err error) {
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加")
		fmt.Println("2. 输入get 表示获取")
		fmt.Println("3. 输入show 表示显示")
		fmt.Println("4. 输入exit 表示推出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("add ok")
			}
		case "get":
			fmt.Println("get")
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("get ok")
			}
			fmt.Printf("从队列中取出一个数：%d\n", val)
		case "show":
			queue.ShowQueue()
			fmt.Println()
		case "exit":
			os.Exit(0)

		}
	}

}
