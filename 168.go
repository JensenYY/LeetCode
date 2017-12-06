package main

func convertToTitle(n int) string {
	dict := []int{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	ret := ""
	var b int
	for i := n; i > 0; i /= 26 {
		i--
		b = i % 26
		ret = string(dict[b]) + ret
	}
	return ret
}
