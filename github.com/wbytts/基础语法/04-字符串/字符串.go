package main

import "fmt"

func main() {
	// 字符串：多个byte的集合，理解为一个字符序列
	// 语法：使用双引号 或 ` `
	// 其中反引号是原样输出，不进行转义，也可以包含换行，即可以是多行字符串

	// 编码：计算机识别0和1，字符是通过编码识别的
	// go 中字符的编码使用的是 utf-8

	/*
	  字符串的常用操作：
	    求字符串的长度：len(s)
	    拼接：+ 或 fmt.Sprintf
	    分割：strings.Split
	    包含：strings.Contains
	    前后缀：strings.HasPrefix，strings.HasSuffix
	    子串位置：strings.Index，strings.LastIndex
	    join操作：strings.Join(a[]string, sep string)
	*/

	var s1 string
	s1 = "wby"
	fmt.Printf("%T, %s\n", s1, s1)

	/*
	  go 中字符有两种：
	    uint8类型：或者叫byte型，代表了 ASCII 码的一个字符
	    rune类型：代表一个 UTF-8 字符串
	*/
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T, %d\n", v1, v1)
	fmt.Printf("%T, %s\n", v2, v2)

	/*
	  字符串是不能直接修改的
	  如果要修改，可以先转换为一个 []rune 或者 []byte，这两种操作都会重新分配内存，并复制字节数组
	*/
	ss1 := "白萝卜"
	ss2 := []rune(ss1) // 把 s1 字符串强制转换成一个rune切片
	ss2[0] = '红'
	fmt.Println(string(ss2))
}
