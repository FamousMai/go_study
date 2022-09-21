package main

import "fmt"
import "runtime"

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println("num = ", num)
}
