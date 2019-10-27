package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var list []int
	inputMap := make(map[int]int)
	for i, val := range nums {
		inputMap[val] = i
		fmt.Println(i, val)
	}
	for i, val := range nums {
		diff := target - val
		index, ok := inputMap[diff]
		if ok {
			list = append(list, i)
			list = append(list, index)
			return list
		}
	}
	return list
}

func main() {
	ar := []int{3, 2, 4}
	fmt.Println(twoSum(ar, 6))
}
