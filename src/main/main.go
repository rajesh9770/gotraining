package main

import (
	"fmt"
	"StringUtils"
)

// return function type
func wrapper() func() int {
	 x:=0
	 return func() int{
	 	x++
	 	return x 
	 }
}

func main() {
//	fmt.Println("Hello world Test!")
//	fmt.Printf("%d %b %o %#x", 42, 42, 42, 42)
//	for i:=1; i<20000; i++ {
//		fmt.Printf("%d %b %o %#x UTF-8 %q\n", i, i, i, i, i)
//	}

	var a, b string = "aa", "bb"
	c :=10
	var f float64
	
	fmt.Printf("%v %v %v %v ", a, b, c, f)
	
	fmt.Printf("%v ", StringUtils.Concat("a", "aa"))
	fmt.Printf("%v ", StringUtils.MyName)
	fmt.Println("")
	
	var x int = 0
	anonymous := func () int{
		x++
		return x
	}
	fmt.Printf("%v ", anonymous())
	fmt.Printf("%v ", anonymous())
	
	fmt.Println("")
	
	anonymous = wrapper()
	//anonymous holds on to x which 0 at this point
	fmt.Printf("%v ", anonymous())
	fmt.Printf("%v ", anonymous())
	
	fmt.Println("")
	var anonymous1 = wrapper()
	//anonymous1 holds on to another copy of x which 0 at this point
	fmt.Printf("%v ", anonymous1())
	fmt.Printf("%v ", anonymous1())
	
	//should be 3
	fmt.Printf("%v ", anonymous())
}



