package main

import (
	"fmt"
	"strconv"
)

var (
	a int = 3
)

func main() {
	//var i int = 0

	/* 	var (
		a int    = 3
		b string = "nice"
	) */
	//i = 3
	fmt.Println(a)
	a := 4
	fmt.Println(a)
	j := 6
	//fmt.Println(i)
	//var k string
	k := strconv.Itoa(j)
	fmt.Printf("%v , %T\n", j, j)
	fmt.Printf("%v , %T\n", k, k)

}
