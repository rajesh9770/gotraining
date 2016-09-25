package main 

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world Test!")
	fmt.Printf("%d %b %o %#x", 42, 42, 42, 42)
	for i:=1; i<20000; i++ {
		fmt.Printf("%d %b %o %#x UTF-8 %q\n", i, i, i, i, i)
	}
}
