package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var ss = "xx"

//b := true 错误用法
var (
	bb = "xx"
	cc = 3
)

func variableZeroValue() {
	fmt.Println("定义变量")
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	fmt.Println("定义变量并初始化")
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	fmt.Println("自动判断变量类型")
	var a, b, c, s = 3, 4, true, "123"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	fmt.Println("使用:=取代var")
	//注意，使用：定义变量只能在函数内，不能在函数外使用
	a, b, c, s := 3, 4, true, "123"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	fmt.Println("hello world")
	//定义变量
	variableZeroValue()
	//定义变量并初始化
	variableInitialValue()
	//自动判断变量类型
	variableTypeDeduction()
	//使用:=取代var
	variableShorter()
	//打印复数的绝对值
	euler()
	//go语言的类型转换是强制的
	triangle()
}
