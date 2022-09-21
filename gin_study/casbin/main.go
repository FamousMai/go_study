package main

import "github.com/casbin/casbin/v2"

func main() {
	e, err := casbin.NewEnforcer("path/to/model.conf", "path/to/policy.csv")

	sub := "alice" // 想要访问资源的用户
	obj := "data1" // 将要被访问的资源
	act := "read"  // 用户对资源实施的操作

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// 处理错误
	}

	if ok == true {
		// 允许 alice 读取 data1
	} else {
		// 拒绝请求，抛出异常
	}

	// 您可以使用 BatchEnforce() 去批量处理一些请求。
	// 这个方法返回一个布尔类型的切片，切片的下标对应二位数组的行标
	// 例如 results[0] 是 {"alice", "data1", "read"} 的结果
	results, err := e.BatchEnforce([][]interface{}{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})
}
