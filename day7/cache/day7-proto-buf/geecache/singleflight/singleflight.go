package singleflight

import "sync"

// call 代表正在进行中，或已经结束的请求。使用 sync.WaitGroup 锁避免重入。
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group 是 singleflight 的主数据结构，管理不同 key 的请求(call)。
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

// Do 方法，接收 2 个参数，第一个参数是 key，第二个参数是一个函数 fn。Do 的作用就是，针对相同的 key，无论 Do 被调用多少次，函数 fn 都只会被调用一次，等待 fn 调用结束了，返回返回值或错误。
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}

// EasyDo
// g.mu 是保护 Group 的成员变量 m 不被并发读写而加上的锁。为了便于理解 Do 函数，我们将 g.mu 暂时去掉。并且把 g.m 延迟初始化的部分去掉，延迟初始化的目的很简单，提高内存使用效率。
func (g *Group) EasyDo(key string, fn func() (interface{}, error)) (interface{}, error) {
	if c, ok := g.m[key]; ok {
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.m[key] = c

	c.val, c.err = fn()
	c.wg.Done()

	delete(g.m, key)

	return c.val, c.err
}
