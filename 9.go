package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	ans := 0
	for i := x; i > 0; i /= 10 {
		ans = ans*10 + (i % 10)
	}
	if ans == x {
		return true
	}
	return false
}
