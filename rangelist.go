package main

import "fmt"

type RangeList struct {
	// TODO: implement
	List [][2]int
}

func (r *RangeList) Add(addedElement [2]int) error {
	// TODO: implement
	if len(r.List) == 0 {
		r.List = append(r.List, addedElement)
		return nil
	}

	var newList [][2]int
	i := 0
	listLen := len(r.List)
	for i < listLen && r.List[i][1] < addedElement[0] {
		newList = append(newList, r.List[i])
		i++
	}

	for i < listLen && r.List[i][0] <= addedElement[1] {
		addedElement[0] = min(r.List[i][0], addedElement[0])
		addedElement[1] = max(r.List[i][1], addedElement[1])
		i++
	}
	newList = append(newList, addedElement)

	if i < listLen {
		newList = append(newList, r.List[i:len(r.List)]...)
	}

	r.List = newList
	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (r *RangeList) Remove(removedElement [2]int) error {
	// TODO: implement
	if len(r.List) == 0 {
		return nil
	}

	var newList [][2]int

	for _, element := range r.List {
		if element[1] < removedElement[0] || element[0] > removedElement[1] {
			newList = append(newList, element)
			continue
		}

		if removedElement[0] > element[0] {
			newList = append(newList, [2]int{element[0], removedElement[0]})
		}
		if removedElement[1] < element[1] {
			newList = append(newList, [2]int{removedElement[1], element[1]})
		}
	}

	r.List = newList
	return nil
}

func (r *RangeList) Print() error {
	// TODO: implement
	fmt.Println(r.List)
	return nil
}

func main() {
	rl := RangeList{}
	rl.Add([2]int{1, 5})
	rl.Print() // Should display: [1, 5)

	rl.Add([2]int{10, 20})
	rl.Print() // Should display: [1, 5) [10, 20)

	rl.Add([2]int{20, 20})
	rl.Print() // Should display: [1, 5) [10, 20)

	rl.Add([2]int{20, 21})
	rl.Print() // Should display: [1, 5) [10, 21)

	rl.Add([2]int{2, 4})
	rl.Print() // Should display: [1, 5) [10, 21)

	rl.Add([2]int{3, 8})
	rl.Print() // Should display: [1, 8) [10, 21)

	rl.Remove([2]int{10, 10})
	rl.Print() // Should display: [1, 8) [10, 21)

	rl.Remove([2]int{10, 11})
	rl.Print() // Should display: [1, 8) [11, 21)

	rl.Remove([2]int{15, 17})
	rl.Print() // Should display: [1, 8) [11, 15) [17, 21)

	rl.Remove([2]int{3, 19})
	rl.Print() // Should display: [1, 3) [19, 21)
}
