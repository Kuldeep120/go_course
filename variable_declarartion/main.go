package main

import (
	"fmt"
	"strconv"
)

// package level variable declaration
var pack string = "package"
// pack = 3232 
// pack:="package"

// variable name declared in package scope with capital case can be access outstide off the package
// lower cass varialbe can access only on the package
// variable length decide the life spane
func main(){
	
	// declarationSyntax()
	// variableCase()
	// variableParsing()
	// numericalTypes()
	// stringLiterals()
	constants()
}

func declarationSyntax(){
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
}

func variableCase(){
	//smaller length variables has lower life spane
	var i int = 2
	fmt.Println(i)

	// large length variables has larger life spane
	var longerRangeVariable string = "121"
	fmt.Println(longerRangeVariable)

	// variable name with Acronyme should be all capitals

	var theUrl string = "3232323" // not recommended
	var theURL string = "3232323" // recommended

	fmt.Println(theUrl)
	fmt.Println(theURL)

}	

func variableParsing(){
	// variable parsing
	var flt float32 = 2111.11
	fmt.Println(flt)

	var inte int
	inte = int(flt)
	var intToFlt float32
	intToFlt = float32(inte) * 11.11
	fmt.Println(inte)
	fmt.Printf("%v, %T\n", inte, inte)
	fmt.Println(intToFlt)
	fmt.Printf("%v, %T\n", intToFlt, intToFlt)


	// number to string conversion

	var str string
	str = string(inte)
	fmt.Printf("%v, %T\n", str, str)

	// strconv package usage for converting numbers to string

	str = strconv.Itoa(inte)
	fmt.Printf("%v, %T", str, str)

	// str = strconv.Itoa(intToFlt)
	// str = strconv.ParseFloat(intToFlt, 32)
	fmt.Printf("%v, %T", str, str)

}

func numericalTypes(){
	// b1 := 1==2
	// b2 := 1==1
	// b3: true

	a := 4232
	b := 32
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// bit operator
	fmt.Println("Bitwise operations\n\n\n")
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	// fmt.Println(a % b)

	// bit shift operator

	a = 10

	fmt.Println(a << 3)
	fmt.Println(a >> 2)

	// flaot literal

	n:= 2.14
	// var n float64 = 12.3
	n = 13.7e72
	n = 2.1E14

	fmt.Println(n)

	// complex numbers
	var ib complex64 = 4- 0i

	ia := 233+332i
	fmt.Println(b)
	fmt.Println(a)
	fmt.Printf("%v, %T\n", real(ia), real(ia))
	fmt.Printf("%v, %T\n", real(ib), real(ib))
}

func stringLiterals(){
	a := "Hello How are you"
	// a[3] = "u"
	a2 := "I am fine"
	fmt.Printf("%v, %T\n", string(a[3]), a[3])
	fmt.Printf("%v, %T\n", a + a2, a + a2)

	a3 := []byte(a)

	fmt.Println(a3)

	// runes

	room := 'a'
	var r rune = 'd'
	// var r2 rune = "sd"

	fmt.Printf("%v, %T\n", room, room)
	fmt.Printf("%v, %T\n", r, r)
	// fmt.Printf("%v, %T\n", r2, r2)

}

func constants(){
	const myConst int = 32
	// const a = 42
	const (
		a = iota + 100
		b
		c
		d
	)
	fmt.Printf("%v, %T\n", myConst, myConst)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	const (
		_ = iota
		e
		f
		g
	)

	fmt.Println(e)

	const (
		_ = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fileSize := 4000000000.
	fmt.Printf("%.2GB", fileSize/GB)
	// fmt.Println(KB)
	// fmt.Println(MB)
	// fmt.Println(GB)
	// fmt.Println(TB)
	// fmt.Println(PB)
	// fmt.Println(EB)
	// fmt.Println(ZB)
	// fmt.Println(YB)

}