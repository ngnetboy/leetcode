package main

import "fmt"

func getMiddle(array []int, low, high int) int {
	temp := array[low]
	for low < high {
		for array[high] >= temp && high > low {
			high -= 1
		}
		array[low] = array[high]

		for array[low] <= temp && high > low {
			low += 1
		}
		array[high] = array[low]
	}
	array[low] = temp
	return low
}

func quickSort(array []int, low, high int) {
	if high <= low {
		return
	}
	middle := getMiddle(array, low, high)
	quickSort(array, low, middle-1)
	quickSort(array, middle+1, high)

}

func main() {
	array := []int{6, 10, 8, 9, 2, 19, 4, 2}
	fmt.Println(array)
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
