package main

func arraySum(x [3]int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func main() {

}
