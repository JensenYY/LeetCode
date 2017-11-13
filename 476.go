package main

import (
	"fmt"
)

func findComplement(num int) int {
    if num == 1{
        return 0
    }
    count := 1
	for count <= num {
		count *= 2
	}
	count -= 1
	return count ^ num
}

func main() {
	fmt.Println(findComplement(5))
}
