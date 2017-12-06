package main

func selfDividingNumbers(left int, right int) []int {
	ret := make([]int, 0, 0)
	for i := left; i <= right; i++ {
		tmp := i
		ok := true
		for tmp > 0 {
			t := tmp % 10
			if t == 0 || i%t != 0 {
				ok = false
			}
			tmp /= 10
		}
		if ok {
			ret = append(ret, i)
		}
	}
	return ret
}
