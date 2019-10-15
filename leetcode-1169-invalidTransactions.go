package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*

如果出现下述两种情况，交易 可能无效：

交易金额超过 ¥1000
或者，它和另一个城市中同名的另一笔交易相隔不超过 60 分钟（包含 60 分钟整）
每个交易字符串 transactions[i] 由一些用逗号分隔的值组成，这些值分别表示交易的名称，时间（以分钟计），金额以及城市。

给你一份交易清单 transactions，返回可能无效的交易列表。你可以按任何顺序返回答案。


示例 1：

输入：transactions = ["alice,20,800,mtv","alice,50,100,beijing"]
输出：["alice,20,800,mtv","alice,50,100,beijing"]
解释：第一笔交易是无效的，因为第二笔交易和它间隔不超过 60 分钟、名称相同且发生在不同的城市。同样，第二笔交易也是无效的。

示例 2：
输入：transactions = ["alice,20,800,mtv","alice,50,1200,mtv"]
输出：["alice,50,1200,mtv"]

示例 3：
输入：transactions = ["alice,20,800,mtv","bob,50,1200,mtv"]
输出：["bob,50,1200,mtv"]


提示：
transactions.length <= 1000
每笔交易 transactions[i] 按 "{name},{time},{amount},{city}" 的格式进行记录
每个交易名称 {name} 和城市 {city} 都由小写英文字母组成，长度在 1 到 10 之间
每个交易时间 {time} 由一些数字组成，表示一个 0 到 1000 之间的整数
每笔交易金额 {amount} 由一些数字组成，表示一个 0 到 2000 之间的整数
输入:
["bob,689,1910,barcelona",
"alex,696,122,bangkok",
"bob,832,1726,barcelona",
"bob,820,596,bangkok",
"chalicefy,217,669,barcelona",
"bob,175,221,amsterdam"]
输出
["bob,689,1910,barcelona","bob,832,1726,barcelona"]
预期结果
["bob,689,1910,barcelona","bob,832,1726,barcelona","bob,820,596,bangkok"]

*/

func invalidTransactions(transactions []string) []string {
	var result []string

	flagMap := make(map[int]bool)
	transActionsLen := len(transactions)
	for i := 0; i < transActionsLen; i++ {
		infoI := strings.Split(transactions[i], ",")
		amountI, _ := strconv.Atoi(infoI[2])
		timeI, _ := strconv.Atoi(infoI[1])
		if amountI > 1000 && !flagMap[i] {
			result = append(result, transactions[i])
			flagMap[i] = true
		}
		for j := i + 1; j < transActionsLen; j++ {
			infoJ := strings.Split(transactions[j], ",")
			amountJ, _ := strconv.Atoi(infoJ[2])
			timeJ, _ := strconv.Atoi(infoJ[1])
			if amountJ > 1000 && !flagMap[j] {
				result = append(result, transactions[j])
				flagMap[j] = true
			}

			if infoI[0] == infoJ[0] && math.Abs(float64(timeI-timeJ)) <= 60 && infoI[3] != infoJ[3] {
				if !flagMap[i] {
					result = append(result, transactions[i])
					flagMap[i] = true
				}

				if !flagMap[j] {
					result = append(result, transactions[j])
					flagMap[j] = true
				}
			}
		}
	}

	return result
}

func main() {
	//s := []string{"bob,689,1910,barcelona", "alex,696,122,bangkok", "bob,832,1726,barcelona", "bob,820,596,bangkok", "chalicefy,217,669,barcelona", "bob,175,221,amsterdam"}
	s1 := []string{"bob,55,173,barcelona", "lee,113,952,zurich", "maybe,115,1973,madrid", "chalicefy,229,283,istanbul", "bob,24,874,shanghai", "alex,568,412,tokyo", "alex,242,1710,milan", "iris,722,879,shenzhen", "chalicefy,281,1586,warsaw", "maybe,246,778,bangkok", "xnova,605,166,newdelhi", "iris,631,991,hongkong", "chalicefy,500,620,tokyo", "chalicefy,380,428,istanbul", "iris,905,180,barcelona", "alex,810,732,shenzhen", "iris,689,389,paris", "xnova,475,298,singapore", "lee,58,709,amsterdam", "xnova,717,546,guangzhou", "maybe,78,435,shenzhen", "maybe,333,145,hongkong", "lee,405,1230,hongkong", "lee,456,1440,tokyo", "chalicefy,286,1071,amsterdam", "alex,55,271,shanghai", "bob,91,273,warsaw", "iris,195,1825,tokyo", "maybe,639,417,madrid", "maybe,305,882,chicago", "lee,443,47,chicago", "chalicefy,958,840,budapest", "lee,162,1239,budapest", "bob,701,505,montreal", "alex,52,1575,munich", "bob,533,1407,amsterdam", "lee,621,491,tokyo", "chalicefy,866,622,rome", "alex,925,455,hongkong", "lee,968,164,moscow", "chalicefy,31,1119,newdelhi", "iris,527,700,warsaw", "bob,286,1694,dubai", "maybe,903,29,barcelona", "maybe,474,1606,prague", "maybe,851,648,beijing", "lee,48,655,chicago", "maybe,378,25,toronto", "lee,922,691,munich", "maybe,411,903,taipei", "lee,651,112,guangzhou", "lee,664,506,dubai", "chalicefy,704,924,milan", "maybe,333,1264,budapest", "chalicefy,587,1112,singapore", "maybe,428,437,moscow", "lee,721,366,newdelhi", "iris,824,1962,beijing", "chalicefy,834,489,istanbul", "alex,639,1473,zurich", "xnova,898,738,tokyo", "chalicefy,585,1313,frankfurt", "xnova,730,759,beijing", "alex,69,892,montreal", "lee,77,91,barcelona", "lee,722,611,taipei", "chalicefy,706,1982,jakarta", "chalicefy,743,584,luxembourg", "xnova,683,322,istanbul", "chalicefy,60,861,prague", "alex,366,871,shenzhen", "chalicefy,77,870,shenzhen", "iris,913,1501,warsaw", "iris,846,1176,warsaw", "bob,873,69,zurich", "alex,601,181,chicago", "chalicefy,118,145,hongkong", "bob,879,982,montreal", "lee,994,950,chicago", "maybe,885,1900,shanghai", "lee,717,1447,shanghai", "chalicefy,71,434,istanbul", "bob,870,968,toronto", "maybe,718,51,beijing", "alex,669,896,istanbul", "chalicefy,639,506,rome", "alex,594,934,frankfurt", "maybe,3,89,jakarta", "xnova,328,1710,rome", "alex,611,571,chicago", "chalicefy,31,458,montreal", "iris,973,696,toronto", "iris,863,148,rome", "chalicefy,926,511,warsaw", "alex,218,1411,zurich", "chalicefy,544,1296,shenzhen", "iris,27,23,montreal", "chalicefy,295,263,prague", "maybe,575,31,munich", "alex,215,174,prague"}
	fmt.Println(invalidTransactions(s1))
}
