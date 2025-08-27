package main

import "fmt"

func sampleSearch(myList []int, myNumber int) int {
	for i := 0; i < len(myList); i++ {
		if myList[i] == myNumber {

			return i
		}
	}

	return -1
}

func main() {
	myList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(sampleSearch(myList, 10))
}
