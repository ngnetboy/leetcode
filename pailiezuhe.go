package main

import "fmt"

func dst2(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if (*used)[i] {
			continue
		}
		p[index] = nums[i]
		(*used)[i] = true
		dst2(nums, index+1, p, res, used)
		(*used)[i] = false
	}
}

func main() {
	num := []int{4, 2, 6}
	p := make([]int, len(num))
	res := [][]int{}
	used := make([]bool, len(num))
	dst2(num, 0, p, &res, &used)
	fmt.Println(res)
}
