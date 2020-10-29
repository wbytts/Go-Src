package main // 表示是main包，可以打包成可执行程序

// 导入语句
import "fmt"

// 函数外部不能写具体的逻辑
// 函数外部只能写一些声明

// 程序的入口函数
func main() {
	// go 的输出语句， 是 fmt 包下的一个方法
	// fmt.Printf("Hello World")

	var n = 100
	// 占位符详解
	fmt.Printf("%T\n", n) // 打印类型
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	var s = "Hello"
	fmt.Printf("04-字符串：%s\n", s)
	fmt.Printf("04-字符串：%v\n", s)
	fmt.Printf("04-字符串：%v\n", s)
}
