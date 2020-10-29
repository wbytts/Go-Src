package main

import "fmt"
import "math"

// 函数外只能放置标识符（变量、常量、函数、类型）的声明

// 程序的入口函数
func main() {
	/*
		基本数据类型：
			布尔类型：bool，表示真假，有 true 和 false 两个值
			数值类型：
				整数：
				小数：
				复数：conplex
			04-字符串：string，go中字符串是基本类型

		复合数据类型：
			array、slice、map、function、pointer、interface、channel
	*/

	// 布尔类型
	var b1 bool
	b1 = true
	fmt.Printf("%T, %t\n", b1, b1)
	b2 := false
	fmt.Printf("%T, %t\n", b2, b2)

	// 整数类型
	// int 是带符号的，uint是不带符号的
	// int 和 uint 的大小由平台决定
	// 也可以使用指定大小的类型：int8、int16、int32、int64、uint8、uint16、uint32、uint64
	// uintptr 无符号整型，用于存储一个指针

	k := int8(123) // 明确指定int8类型，否则默认为int类型
	fmt.Println(k)

	// 内建函数 len(x) 获取长度
	fmt.Println(len("asd"))

	// 浮点类型 float
	// float32、float64
	// go 中的小数，默认都是 float32 类型
	// float64 类型需要显式声明
	fmt.Println("浮点数的最大值（32/64）：", math.MaxFloat32, math.MaxFloat64)

	// 复数类型 complex
	// complex64、complex128

	/*
		类型转换语法：Type(Value)

		注意：兼容类型可以转换

		常数：在有需要的时候，自动转型
		变量：需要手动转型
	*/

	var a int8
	a = 3
	fmt.Printf("%T, %d\n", a, a)
	var b int16
	b = int16(a)
	fmt.Printf("%T, %d\n", b, b)
}
