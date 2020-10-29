package main // 表示是main包，可以打包成可执行程序

// 导入语句
import "fmt"

// 函数外部不能写具体的逻辑
// 函数外部只能写一些声明

// 程序的入口函数
func main() {
	// go 的输出语句， 是 fmt 包下的一个方法
	fmt.Println("Hello World")
}
