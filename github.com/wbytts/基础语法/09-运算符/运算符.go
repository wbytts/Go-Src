package main

import "fmt"

func main() {
	// 算术运算符
	/*
			加减乘除  + - * /
			% 求余

			注：++ 、 -- 在go里面是单独的语句，并不是运算符
		    即 ++、-- 不能参与表达式的组成！
	*/

	a := 10
	b := 3

	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	sub := a - b
	fmt.Printf("%d - %d = %d\n", a, b, sub)

	mul := a * b
	fmt.Printf("%d * %d = %d\n", a, b, mul)

	div := a / b
	fmt.Printf("%d / %d = %d\n", a, b, div)

	mod := a % b
	fmt.Printf("%d %% %d = %d\n", a, b, mod)

	c := 5
	fmt.Printf("c = %d\n", c)
	c++ // ++ 和 -- 不支持参与运算，只能单独使用
	fmt.Printf("c = %d\n", c)

	// 关系运算符
	// > >= < <= == !=
	// 关系运算符的结果是布尔类型

	// 逻辑运算符
	/*
		&&：逻辑与
		||：逻辑或
		!：逻辑非
	*/

	/*
		位运算符：
			&   按位与
			|   按位或
			^	按位异或（二元）、按位取反（一元（前缀））
			&^  位清空（为0取值，非零取0）
			<<  左移
			>>  右移
	*/

	/*
		赋值运算符 与 复合赋值运算符
			......
	*/

	ba := 60
	bb := 11
	fmt.Printf("a: %d, %b\n", ba, ba)
	fmt.Printf("b: %d, %b\n", bb, bb)

}
