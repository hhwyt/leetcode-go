package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, n := range nums {
		_, prs := m[target-n]
		if prs {
			return []int{m[target-n], i}
		} else {
			m[n] = i
		}
	}

	return nil
}

func main() {
	fmt.Println("vim-go")
}
