package main

import "fmt"

func main() {

	// map[键类型]值类型

	m1 := map[string]int{"a": 3, "b": 4}
	fmt.Println(m1)

	// 基本使用
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)

	// 判断某个键是否存在的写法
	// value, ok := map[key]

	fmt.Println("----------")

	// 使用 for range 遍历map
	// 注意：遍历map时的元素顺序与添加键值对的顺序无关
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	fmt.Println("----------")

	// 只遍历 key
	for k := range scoreMap {
		fmt.Println(k)
	}

	fmt.Println("----------")

	// 删除一组键值对
	delete(scoreMap, "小明")
	fmt.Println(scoreMap)

	fmt.Println("----------")
}
