package main

func numJewelsInStones(j string, s string) int {
	result := make(map[int32]int)
	num := 0

	for _, value := range s {
		if _, ok := result[value]; ok {
			result[value] += 1
			continue
		}
		result[value] = 1
	}

	for _, value := range j {
		if _, ok := result[value]; ok {
			num += result[value]
		}
	}
	return num
}

// func main() {
// 	j := "aA"
// 	s := "aAAbbbb"
// 	fmt.Println(numJewelsInStones(j, s))
// }
