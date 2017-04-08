package main

import "fmt"

type Test struct {
	num  int
	desc string
}

// main: create and print a slice of many types
func main() {
	var i = make([]interface{}, 1)
	i[0] = "interface{} slice of length 1"
	f := 3.14
	s := "Hello world"
	i = append(i, f, s, Test{1, "test01"})
	fmt.Println(i)
	for idx, v := range i {
		fmt.Printf("%v :: %#v :: %T\n", idx, v, v)
	}
}
