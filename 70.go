package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	arr := make([]int, n+1, n+1)
	arr[1] = 1
	arr[2] = 2
	i := 3
	for ; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[i-1]
}

func main() {
	fmt.Println(climbStairs(4))
}
