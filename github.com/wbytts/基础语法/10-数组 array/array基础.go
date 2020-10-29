package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	/*
	   数组：存放元素的容器
	     必须指定存放的元素的类型和容量（长度）
	     长度是类型的一部分
	*/
	var arr2 [5]int // 数组中默认元素都是零值
	fmt.Println(arr2)

	var arr3 = [...]int{1, 2, 3, 4, 5} // 三个点，表示根据初始值自动推算出元素个数
	fmt.Println(arr3)

	var arr4 = [5]int{1, 2} // 初始化不写全的话，剩下的默认是0
	fmt.Println(arr4)

	arr5 := [5]int{5, 4, 3, 2, 1}
	fmt.Println(arr5)

	// 多维数组的用法
	arr6 := [3][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}, [3]int{7, 8, 9}}
	fmt.Println(arr6)

	// 数组是值类型！
}
