package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		math/rand
		伪随机数, 根据一定的算法公式算出来的
	*/
	num1 := rand.Int()
	fmt.Println(num1)

	// 设置随机数种子，让每次运行都是不一样的随机数
	rand.Seed(time.Now().UnixNano())
	num2 := rand.Int()
	fmt.Println(num2)
}
