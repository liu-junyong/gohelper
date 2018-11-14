package gohelper

import (
	"fmt"
	"strconv"
	"strings"
)

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

func StringToSlice(s string) []int32 {
	ret := make([]int32, 0)
	vs := strings.Split(s, ",")
	for _, v := range vs {
		r, _ := strconv.Atoi(v)
		ret = append(ret, int32(r))
	}
	return ret
}
