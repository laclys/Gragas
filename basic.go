package main

// import (
// 	"fmt"
// 	"math"
// 	"math/cmplx"
// )

// var (
// 	aa  = 3
// 	sss = "kk"
// )

// func variableZeroValue() {
// 	var a int
// 	var b string
// 	fmt.Printf("%d %q\n", a, b)
// }

// func variableInitialValue() {
// 	var a, b int = 3, 4
// 	fmt.Printf("%d %d\n", a, b)
// }

// func variableTypeDeduction() {
// 	var a, b, c, s = 3, 4, true, "def"
// 	fmt.Println(a, b, c, s)
// }

// func variableShorter() {
// 	a, b, c, s := 3, 4, true, "def"
// 	fmt.Println(a, b, c, s)
// }

// func triangle() {
// 	var a, b int = 3, 4
// 	var c int
// 	c = int(math.Sqrt(float64(a*a + b*b)))
// 	fmt.Println(c)
// }

// func euler() {
// 	fmt.Println(
// 		cmplx.Pow(math.E, 1i*math.Pi) + 1)
// }

// func consts() {
// 	const (
// 		filename = "abc.txt"
// 		a, b     = 3, 4
// 	)
// 	var c int
// 	c = int(math.Sqrt(a*a + b*b))
// 	fmt.Println(c)
// }

// func enums() {
// 	// 枚举类型
// 	const (
// 		cpp = iota
// 		java
// 		python
// 		golang
// 	)
// 	// b, kb, sb, tb
// 	const(
// 		b = 1 << (10 * iota)
// 		kb
// 		sb
// 		tb
// 	)

// 	fmt.Println(cpp, java, python, golang)
// 	fmt.Println(b, kb, sb, tb)
// }

// func main() {
// 	fmt.Println("hello world")
// 	variableZeroValue()
// 	variableInitialValue()
// 	variableTypeDeduction()
// 	variableShorter()
// 	fmt.Println(aa, sss)
// 	euler()
// 	triangle()
// 	consts()
// 	enums()
// }
