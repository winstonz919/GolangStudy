package main

import (
	"fmt"
	"reflect"
)

func Foo(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t)
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println(x)
	fmt.Println("non-empty interface")
}

func main() {
	var x *int = nil
	Foo(x)

	const a = 1
	var b = 1
	fmt.Println(a == b)
}
