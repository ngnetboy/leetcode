package main

import (
	"fmt"
	"math"
)

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

*/

func reverse(x int) int {
	if math.Abs(float64(x)) < 10 {
		return x
	}

	result := 0
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result > 2147483647 || result < -2147483647 {
		return 0
	}
	return result
}

func main() {
	//fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(1534236469))
	fmt.Printf("%f", math.Pow(2, 31)-1)

}
