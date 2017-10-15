package main

import (
	"fmt"
)

/*
{ 'I' , 1 },
{ 'V' , 5 },
{ 'X' , 10 },
{ 'L' , 50 },
{ 'C' , 100 },
{ 'D' , 500 },
{ 'M' , 1000 }
*/

func romanToInt(s string) int {
	//In golang 'rune' means 'char',but if want to print it must convert to 'string'
	rmap := make(map[string]int)
	rmap["I"] = 1
	rmap["V"] = 5
	rmap["X"] = 10
	rmap["L"] = 50
	rmap["C"] = 100
	rmap["D"] = 500
	rmap["M"] = 1000

	sum := 0
	for i := 0; i < len(s)-1; i++ {
		if rmap[string(s[i])] < rmap[string(s[i+1])] {
			sum -= rmap[string(s[i])]
		} else {
			sum += rmap[string(s[i])]
		}
	}

	return sum + rmap[string(s[len(s)-1])]
}

func main() {
	i := romanToInt("DCXXI")
	fmt.Println(i)
}
