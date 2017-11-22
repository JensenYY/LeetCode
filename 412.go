package main

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	if n <= 0 {
		return nil
	}
	//make a slice the len is n,the capacity is n
	ret := make([]string, n, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 != 0 {
			ret[i-1] = "Fizz"
		} else if i%3 != 0 && i%5 == 0 {
			ret[i-1] = "Buzz"
		} else if i%3 == 0 && i%5 == 0 {
			ret[i-1] = "FizzBuzz"
		} else {
			ret[i-1] = strconv.Itoa(i)
		}
	}
	return ret
}
