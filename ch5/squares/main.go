// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 135.

// The squares program demonstrates a function value with state.
package main

import "fmt"

//!+
// squares returns a function that returns squares函数返回一个函数，后者包含下一次要用到数的平方
// the next square number each time it is called.
//函数变量不仅是代码，还可以拥有状态。go的函数变量就是闭包
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}

//!-
