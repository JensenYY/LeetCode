package main

import ()

func reverseString(s string) string {
	//sb mean's stringBuffer
	sb := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		sb[i], sb[j] = sb[j], sb[i]
		i++
		j--
	}
	return string(sb)
}
