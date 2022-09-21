package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

// 添加
func InsertCatNode(head *CatNode, newCatNode *CatNode) {

	// 判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //构成一个环形
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}

	// 定义一个临时的变量，帮忙的,找到环形最后的结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	//加入到链表中
	temp.next = newCatNode
	newCatNode.next = head

}

// 输出这个环形的链表
func ListCircleLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也")
		return
	}

	for {
		fmt.Printf("猫的信息为=[id=%d name=%s] -> \n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

// 删除
func DeleteCircleLink(head *CatNode, id int) {
	temp := head
	helper := head

	// 空链表
	if temp.next == nil {
		fmt.Printf("del fail Link empty")
		return
	}

	// 如果只有一个结点
	if temp.next == head {
		temp.next = nil
		return
	}

	// 如果有两个包含两个以上结点
	for {
		if temp.next == head { // 如果到这来，说明我比较到最后一个（最后一个还没比较）
			break
		}

		if temp.no == id {
			// 恭喜找到，我们也可以在直接删除
		}

		temp = temp.next     // 移动【比较】
		helper = helper.next //移动【借助这个来删除】
	}
}

func main() {
	fmt.Println("环形链表")

	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	InsertCatNode(head, cat1)
	ListCircleLink(head)
}
