package main

import "fmt"

func main() {

	age := 24

	/*
	  和其他语言的 switch 不同的地方
	    默认每个case执行完了就会结束，不会接着执行之后的case，也就是说，不需要像其他语言一样加 break
	*/

	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

}
