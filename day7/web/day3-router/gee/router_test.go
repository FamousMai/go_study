package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok1 := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok2 := reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok3 := reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})

	if !ok1 && !ok2 && !ok3 {
		t.Fatal("test parsePattern failed", ok1, ok2, ok3)
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/geektutu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("shouldn't match /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu")
	}
	fmt.Printf("matched path : %s, params['name']:%s\n", n.pattern, ps["name"])
}
