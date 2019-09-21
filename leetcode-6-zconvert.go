package main

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

*/

import (
	"fmt"
)

func convert(s string, numRows int) string {
	lenS := len(s)
	if lenS == 0 || lenS == 1 || numRows == 1 {
		return s
	}
	sum := numRows*2 - 2
	var result string
	for i := 0; i < numRows; i++ {
		for j := i; j < lenS; j += sum {
			result = fmt.Sprintf("%s%c", result, s[j])
			skipLen := j - i*2 + sum
			if i != 0 && i != numRows-1 && skipLen < lenS {
				result = fmt.Sprintf("%s%c", result, s[skipLen])
			}
		}
	}
	return result
}

func main() {
	res := convert("A", 1)
	fmt.Println(res)
}
