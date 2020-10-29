package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello World")
	}

	// break 跳出整个循环
	// continue 结束本次循环

	// for range 用法
	// 用于循环一个序列
	s := "HelloWorld"
	for i, v := range s { // 第一个是索引，第二个是值
		fmt.Printf("%d %c\n", i, v)
	}

	// 死循环，什么都不写就是死循环
	//for {
	//	fmt.Println("Hello......")
	//}

}
