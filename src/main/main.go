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
	defineSliceEx1()
	defineSliceEx2()
	defineMapEx1()
	
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
	
	// if you do not call make, then it uses the previous record
	record1[0] = "D"
	record1[1] = "Surname"
	record1[2] = "1/2/1957"
	trans[3] = record1
	
	for i:=0; i<len(trans); i++{
		fmt.Println(trans[i])
	}
	fmt.Println(trans)
}

func defineSliceEx1(){
	student := [] string{} // no len no capcity, but initilized
	students := [][] string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil) //false, b/c it's initialized
	fmt.Println(students == nil)
	student = append(student, "test") // works
	fmt.Println(student)
}

func defineSliceEx2(){
	var student [] string  // just the definition, but no initialization
	var students [] string
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil) // true b/c it's not inititialized
	fmt.Println(students == nil)
	student = append(student, "test2") // works, still can do append, append bring it to life
	fmt.Println(student)
	fmt.Println(student == nil) // false b/c it's now inititialized
}

func defineMapEx1(){
	// there is no append function, this is uselless, you can't add values to map in this case
	var myMap map[string]string
	fmt.Println(myMap == nil) //true
	fmt.Println(myMap)
	//myMap["Key1"] = "value1"
	//fmt.Println(myMap)
	
	var myMap2 = map[string]string{}
	fmt.Println(myMap2 == nil) //false
	fmt.Println(myMap2)
	myMap2["Key1"] = "value1"
	fmt.Println(myMap2)
	
	
	var myMap3 = make(map[string]string)
	fmt.Println(myMap3 == nil) //false
	fmt.Println(myMap3)
	myMap3["Key2"] = "value2"
	fmt.Println(myMap3)
	
	var myMap4 = map[string]string{"Name" : "Rajesh", "Address" : "harwood road"}
	fmt.Println(myMap4 == nil) //false
	fmt.Println(myMap4)
	myMap4["Key1"] = "value1"
	fmt.Println(myMap4)
	
}
