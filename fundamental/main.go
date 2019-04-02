package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	//declaration
	var a bool
	var b float64
	var c int
	var d string
	var e complex128
	// declaration with initialization
	var aa bool = true
	//multiple
	var bb, cc int32 = 20, 30
	//without data type
	var dd = 40
	//short declare n init without datatype
	ee := 50
	//must be initialized variabel n never change
	const ff = 15

	//initialization -> default 0, empty, false
	a = true
	b = 2.2
	c = 5
	d = "Testing"
	e = cmplx.Acos(-2 + 12i)

	//print
	fmt.Println("Value a: ", a)
	fmt.Println("Value b: ", b)
	fmt.Println("Value c: ", c)
	fmt.Println("Value d: ", d)
	fmt.Println("Value e: ", e)
	fmt.Println("Value aa: ", aa)
	fmt.Println("Value bb: ", bb)
	fmt.Println("Value cc: ", cc)
	fmt.Println("Value dd: ", dd)
	fmt.Println("Value dd: ", ee)
	fmt.Println("Value ff: ", ff)

	//convert
	bb = int32(math.Sqrt(b))
	fmt.Println("Value convert b: ", bb)

}
