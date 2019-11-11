package main

import "fmt"

//1
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

//2
func twoSum2(nums []int, target int) []int {

	tmpMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		subv := target - nums[i]
		if k, ok := tmpMap[subv]; ok {
			return []int{k, i}
		}
		tmpMap[nums[i]] = i
	}

	return []int{}
}

func main() {

	nums := []int{3, 2, 4}
	target := 6
	result := twoSum2(nums, target)
	fmt.Printf("%v\n", result)
}
