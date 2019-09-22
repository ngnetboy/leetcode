package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/

//
//func checkEque(a []int, b []int) bool {
//	for _, value := range a {
//
//	}
//}

func threeSum(nums []int) [][]int {
	result := [][]int{}

	a := []int{1, 0, -1}
	b := []int{-1, 0, 1}
	sort.Ints(a)
	sort.Ints(b)
	fmt.Println(reflect.DeepEqual(a, b))

	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
