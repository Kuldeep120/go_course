package main

import (
	"fmt"
)

// package level variable declaration
var pack string = "package"
// pack = 3232 
// pack:="package"

// variable name declared in package scope with capital case can be access outstide off the package
// lower cass varialbe can access only on the package
// variable length decide the life spane
func main(){
	// variable declaration assignment syntax 1
	var var1 int
	var1 = 12
	fmt.Println(var1)
	fmt.Printf("%v\n", var1)
	
	// variable declaration and assgnment syntax 2
	var var2 int = 23
	fmt.Println(var2)

	// variable declaration and assignment syntax 3
	var3 := 32
	pack := 32
	fmt.Println(var3)

	// var parsint float32 = "232"
	// fmt.Println(parsint)
	fmt.Println(pack)


	//smaller length variables has lower life spane
	var i int = 2
	fmt.Println(i)

	// large length variables has larger life spane
	var longerRangeVariable string = "121"
	fmt.Println(longerRangeVariable)

	// variable name with Acronyme should be all capitals

	// var theUrl string = "3232323" // not recommended
	// var theURL string = "3232323" // recommended

	// variable parsing

	var flt float32 = 2111.11
	fmt.Println(flt)

	var inte int
	inte = int(flt)
	fmt.Println(inte)
}