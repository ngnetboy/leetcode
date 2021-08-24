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

func test1() {
	num := []int{4, 2, 6}
	p := make([]int, len(num))
	res := [][]int{}
	used := make([]bool, len(num))
	dst2(num, 0, p, &res, &used)
	fmt.Println(res)
}

func permutation(S string) []string {
	if len(S) == 1 {
		return []string{S}
	}
	// 与拼接得到的各个字符串再进行拼接
	ret := []string{}
	for i, s := range S {
		// 差了第i个字符的剩余字符串往下传，并将得到的结果进行合并
		tmp := fmt.Sprintf("%s%s", S[:i], S[i+1:])
		fmt.Println(tmp, i, S, S[:i], S[i+1:])
		res := permutation(tmp)
		for _, r := range res {
			ret = append(ret, fmt.Sprintf("%c%s", s, r))
		}
	}
	return ret
}

func main() {
	fmt.Println(permutation("123"))
}
