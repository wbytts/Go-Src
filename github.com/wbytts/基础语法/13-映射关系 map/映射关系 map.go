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

	// 元素为map类型的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	fmt.Println("----------")

	// 值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

	fmt.Println("----------")
}
