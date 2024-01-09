package main

import (
	"fmt"
)

/*

在给定的网格中，每个单元格可以有以下三个值之一：

值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。

返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。

示例 1：

输入：[[2,1,1],[1,1,0],[0,1,1]]
输出：4
示例 2：

输入：[[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
示例 3：

输入：[[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
提示：

1 <= grid.length <= 10
1 <= grid[0].length <= 10
grid[i][j] 仅为 0、1 或 2

*/

func orangesRotting(grid [][]int) int {
	result := 0
	gridLen := len(grid)
	for {
		flag := 0
		badFruit := [][]int{}

		//get the index of bad fruit
		for i := 0; i < gridLen; i++ {
			gridILen := len(grid[i])
			for j := 0; j < gridILen; j++ {
				if grid[i][j] == 2 {
					badFruit = append(badFruit, []int{i, j})
				}
			}
		}

		// deal the bad fruit
		for _, value := range badFruit {
			i, j := value[0], value[1]
			gridILen := len(grid[i])
			// left
			if j-1 >= 0 && grid[i][j-1] == 1 {
				grid[i][j-1] = 2
				flag = 1
			}
			//right
			if j+1 < gridILen && grid[i][j+1] == 1 {
				grid[i][j+1] = 2
				flag = 1
			}
			//up
			if i-1 >= 0 && grid[i-1][j] == 1 {
				grid[i-1][j] = 2
				flag = 1
			}
			//down
			if i+1 < gridLen && grid[i+1][j] == 1 {
				grid[i+1][j] = 2
				flag = 1
			}
		}

		if flag != 1 {
			break
		}
		result += 1
	}

	//check fruit
	for i := 0; i < gridLen; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return result
}

func testCopy(s [][]int) {
	temp := [][]int{}
	temp = append(temp, []int{1})
	temp[0] = append(temp[0], 2)
	//copy(temp, s)

	fmt.Println(temp, s)
	//temp[0][0] = 100
	//fmt.Println(temp, s)
}

// func main() {
// 	//s1 := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
// 	//s2 := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
// 	s3 := [][]int{{0, 2}}
// 	fmt.Println(orangesRotting(s3))
// 	//testCopy(s1)
// }
