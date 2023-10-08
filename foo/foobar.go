package foo

import (
	"strconv"
)

func Filtnum(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "Foobar"
	} else if num%5 == 0 {
		return "Bar"
	} else if num%3 == 0 {
		return "Foo"
	} else {
		return strconv.Itoa(num)
	}
}

func Foobar(num int) string {
	result := ""
	for i := 1; i <= num; i++ {
		if i == num {
			result += Filtnum(i)
		} else {
			result += Filtnum(i) + "->"
		}
	}
	return result
}
