package main

import (
	"fmt"
)

func reverseString(s []byte) {
	lenS := len(s)
	fmt.Println(lenS)
	for i := 0; i < lenS/2; i++ {
		s[i], s[lenS-1-i] = s[lenS-1-i], s[i]
	}
}

// func main() {
// 	s := []byte("hell")
// 	reverseString(s)
// 	fmt.Println(string(s))
// }
