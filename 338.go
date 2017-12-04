package main

import "fmt"

func countBits(num int) []int {
	if num <= 0 {
		return nil
	}
	ret := make([]int, num+1, num+1)
	for i := 0; i <= num/2; i++ {
		ret[2*i] = ret[i]
		if 2*i+1 <= num {
			ret[2*i+1] = ret[i] + 1
		}
	}
	return ret
}

func main() {
	fmt.Println(countBits(4))
}
