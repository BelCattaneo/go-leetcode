package main

import "fmt"

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for i, num := range nums {
		for key, value := range sumMap {
			if value + num == target {
				return []int{key, i}
			}
		}
		sumMap[i] = num
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
}
