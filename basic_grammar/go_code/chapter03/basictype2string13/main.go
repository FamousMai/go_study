package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num3 int = 99
	// var num4 float64 = 23.456

	str := strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str = %q\n", str, str)
}
