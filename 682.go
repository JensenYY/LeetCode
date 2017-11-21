package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func calPoints(ops []string) int {
	count := 0
	randRecord := list.New()
	for _, v := range ops {
		if 0 == strings.Compare("C", v) {
			//cancel the head score
			score := randRecord.Front()
			randRecord.Remove(score)
			count -= score.Value.(int)
			continue
		}
		if 0 == strings.Compare("D", v) {
			//add the double head score
			score := randRecord.Front()
			s := score.Value.(int)
			s *= 2
			randRecord.PushFront(s)
			count += s
			continue
		}
		if 0 == strings.Compare("+", v) {
			//add the head and head^ score
			head := randRecord.Front()
			s1 := head.Value.(int)
			s2 := head.Next().Value.(int)
			randRecord.PushFront(s1 + s2)
			count += (s1 + s2)
			continue
		}
		s, _ := strconv.Atoi(v)
		count += s
		randRecord.PushFront(s)
	}
	return count
}

func main() {
	score := []string{"5","-2","4","C","D","9","+","+"}
	fmt.Println(calPoints(score))
}
