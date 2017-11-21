package main

import (

)

type MyCalendar struct {
	booingField map[int]int
}

func Constructor() MyCalendar {
	mc := MyCalendar{
		booingField: make(map[int]int),
	}
	return mc
}

func (this *MyCalendar) Book(start int, end int) bool {
	if start > end {
		return false
	}
	for k, v := range this.booingField {
		if start < k && end > k {
			return false
		}
		if start < v && end > v {
			return false
		}
		if start >= k && end <= v {
			return false
		}
	}

	this.booingField[start] = end
	return true
}
