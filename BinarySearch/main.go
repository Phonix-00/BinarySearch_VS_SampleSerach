package main

import "fmt"

func binarySearch(arr []int, myNumber int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == myNumber {

			return mid
		}

		if guess > myNumber {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	myList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(binarySearch(myList, 10))
}
