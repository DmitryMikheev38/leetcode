package main

import (
	"fmt"
	"sort"
)

var nums []int = []int{3, 3}

func main() {

	res := twoSum(nums, 6)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {

	mumsM := map[int][]int{}

	for idx, num := range nums {
		mumsM[num] = append(mumsM[num], idx)
	}

	sort.Ints(nums)

	i := 0
	j := len(nums) - 1

	for i < j {
		if nums[i]+nums[j] == target {
			break
		}
		if nums[i]+nums[j] < target {
			i++
		}
		if nums[i]+nums[j] > target {
			j--
		}
	}
	res := []int{mumsM[nums[i]][0], mumsM[nums[j]][len(mumsM[nums[j]])-1]}

	return res
}
