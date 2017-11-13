package main

import (
	"fmt"
)

//too slow
func reverseWords(s string) string {
	var str string

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			for j := i - 1; j >= 0 && s[j] != ' '; j-- {
				str += string(s[j])
			}
			str += " "
		}
		if i == len(s)-1 {
			for j := i; j >= 0 && s[j] != ' '; j-- {
				str += string(s[j])
			}
		}
	}

	return str
}

/*Just 12s
func reverseWords(s string) string {
    b := []byte(s)
    start := 0

    for i := 0; i < len(b); i++ {
        if b[i] == ' ' {
            reverseWord(b, start, i-1)
            start = i + 1
        }
    }
    reverseWord(b, start, len(b) - 1)
    return string(b)
}

func reverseWord(b []byte, start, end int){
    for start <= end {
        b[start],b[end] = b[end],b[start]
        start += 1
        end -= 1
    }
}
*/

func main() {
	s := "Hello World! Boy"
	fmt.Println(reverseWords(s))
}
