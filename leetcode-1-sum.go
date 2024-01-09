package main

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func towSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func towSum2(nums []int, target int) []int {
	mapIndex := make(map[int]int, len(nums))

	for index, val := range nums {
		subVal := target - val
		if mapI, ok := mapIndex[subVal]; ok {
			return []int{mapI, index}
		}
		mapIndex[val] = index
	}
	return []int{}
}

// func main() {
// 	val := []int{2, 7, 11, 15}
// 	result := towSum2(val, 9)
// 	fmt.Println(result)
// }
