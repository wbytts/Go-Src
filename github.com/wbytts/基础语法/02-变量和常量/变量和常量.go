package main

// 导入语句
import "fmt"

// 函数外只能放置标识符（变量、常量、函数、类型）的声明

// 程序的入口函数
func main() {
	// go 里面的变量需要先声明再使用
	// var 变量名 变量类型
	var name string

	// 变量声明了之后，默认是这种类型的空值
	// 声明变量之后，可以更改变量的值
	name = "wby"

	// 在 go 中，声明了局部变量之后，必须使用，不然编译不过去

	// 一次声明多个变量
	var (
		a string
		b int
		c bool
		d float32
	)

	a = "1"
	b = 2
	c = true
	d = 1.2

	// 变量的初始化
	// var 变量名 类型 = 表达式
	var age int = 18
	// 也可以一次初始化多个变量
	// var name, age = "Q1mi", 20

	// 类型推导
	// 有时候我们会将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型，完成初始化
	var mm = 123
	fmt.Print(mm)

	// 短变量声明
	// 在函数内部，可以使用更简略的 := 方式声明并初始化变量
	sum := 123
	aa, bb, cc := 11, 22, 33
	fmt.Print(aa, bb, cc)

	// 常量使用const定义，常量的数值不能改变
	const ca = 1

	// 使用常量组作为枚举，一组相关的数值
	const (
		SPRING  = 0
		SUMMERY = 1
		AUTUMN  = 2
		WINTER  = 3
	)

	// 批量声明常量时，如果只声明了第一个常量，则后面的与前面的值一样
	const (
		a1 = 100
		a2
		a3
	)
	// iota，常量计数器， 从0开始，依次递增
	// 即，每新增一个常量的声明， iota 的值就会增加 1
	// iota 在 const 关键字出现的时候，将会被重置为 0
	// 在const定义中使用，可以实现类似枚举的功能
	const (
		s1 = iota
		s2
		s3
	)

	// 应用举例，定义存储级别
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

	fmt.Println(name)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(age)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(sum)

	// 打印变量的地址
	fmt.Printf("sum的地址:%p\n", &sum)
	sum = 3123 // 改变sum的数值，内存地址不改变
	fmt.Printf("sum的地址:%p\n", &sum)

}
