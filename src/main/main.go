package main

import (
	"fmt"
	//"StringUtils"
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

	exec1()
	exec2()
	exec3()
	sliceEx()
	sliceEx2()
	sliceEx3()
	
	/**
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
	**/
}


func isEven(a int) (int, bool){
	return a/2, a%2 ==1;
}


func exec1(){
	a, even := isEven(20)
	fmt.Println(a, even)
	a, even = isEven(21)
	fmt.Println(a, even)	
}
func exec2(){
	
	isEven := func(a int) (int, bool) {return a/2, a%2==1}
	a, even := isEven(20)
	fmt.Println(a, even);	
}

func exec3(){
	
	maxx := max(1,2,3,4,5)
	
	fmt.Println(maxx);	
}

func max (a ...int) int{
	var ret int
	for _, it := range a{
		if ret < it {
			ret = it 
		}
	}
	return ret
}

func sliceEx(){
	mySlice := make ( [] int, 5, 5)
	
	fmt.Println(len(mySlice) , " " , cap(mySlice))
	
	for i:=0; i<80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println(i, " " , len(mySlice) , " " , cap(mySlice))
	}

}

func sliceEx2(){
	mySlice1 := [] int {1,2,3}
	mySlice2 := [] int {3,4,5}
	
	mySlice := append (mySlice1, mySlice2...)
	fmt.Println(mySlice)
	
}

func sliceEx3(){
	trans := make([][] string, 4, 5)
	
	record1 := make([] string, 4, 5)
	record1[0] = "A"
	record1[1] = "Surname"
	record1[2] = "1/2/1945"
	trans[0] = record1
	
	record1 = make([] string, 4, 5)
	record1[0] = "B"
	record1[1] = "Surname"
	record1[2] = "1/2/1947"
	trans[1] = record1
	
	record1 = make([] string, 4, 5)
	record1[0] = "C"
	record1[1] = "Surname"
	record1[2] = "1/2/1967"
	trans[2] = record1
	
	record1[0] = "D"
	record1[1] = "Surname"
	record1[2] = "1/2/1957"
	trans[3] = record1
	
	for i:=0; i<len(trans); i++{
		fmt.Println(trans[i])
	}
	fmt.Print(trans)
}