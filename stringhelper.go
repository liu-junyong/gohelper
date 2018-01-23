package gohelper

import "fmt"

func SliceToString(s *[]int32) string {
	ret := ""
	for i := range *s {
		if i == 0 {
			ret += fmt.Sprintf("%d", (*s)[i])

		} else {
			ret += ","
			ret += fmt.Sprintf("%d", (*s)[i])
		}
	}
	return ret
}
