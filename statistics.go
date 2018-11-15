package gohelper

import (
	"github.com/liu-junyong/go-logger/logger"
	"sort"
)

func Max(first int32, nums ...int32) int32 {
	max := first
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func Min(first int32, nums ...int32) int32 {
	min := first
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func Min_pos(nums ...int32) int32 {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	min := nums[0]
	minpos := 0
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
			minpos = i
		}
	}
	return int32(minpos)
}

//
//func Count_Card(Cards *[]int32, card int32) int {
//	ret := 0
//	for i := range *Cards {
//		if (*Cards)[i] == int32(card) {
//			ret++
//		}
//	}
//	return ret
//}

func Count(card int32, Cards ...int32) int32 {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	ret := int32(0)
	for i := range Cards {
		if (Cards)[i] == int32(card) {
			ret++
		}
	}
	return ret
}

//func Count_if(Cards ...int32,func(i, j int) bool { return nums[i] < nums[j] }) int32 {
//func Count_if(Cards ...int32,func(i int) bool { return nums[i] < nums[j] }) int32 {
func Count_if(cards *[]int32, compare func(int32) bool) int32 {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	ret := int32(0)
	for i := range *cards {
		if compare(i) {
			ret++
		}
	}
	return ret
}

func Most(nums ...int32) (int32, int32) {
	mostct := int32(0)
	mostval := int32(0)
	for _, num := range nums {
		ct := Count(num, nums...)
		if ct >= mostct {
			mostct = ct
			mostval = num
		}
	}
	return mostval, mostct
}

func Least(nums ...int32) int32 {
	least := int32(0)
	leastVal := int32(0)
	for _, num := range nums {
		ct := Count(num, nums...)
		if ct <= least {
			least = ct
			leastVal = num
		}
	}
	return leastVal
}

//判断连续
func Series(nums ...int32) bool {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 != nums[i+1] {
			return false
		}
	}
	return true
}

//是否一致
func Same(nums ...int32) bool {
	for i := 1; i < len(nums); i++ {
		if nums[0] != nums[i] {
			return false
		}
	}
	return true
}
