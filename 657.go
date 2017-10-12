package main

import (
	"fmt"
)

func judeCircle(moves string) bool {
	x := 0
	y := 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'U':
			y++
			break
		case 'D':
			y--
			break
		case 'L':
			x--
			break
		case 'R':
			x++
			break
		}
	}
	if 0 == x && 0 == y {
		return true
	}
	return false
}

func main() {
	if judeCircle("LR") {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
