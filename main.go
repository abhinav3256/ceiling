package main

import "fmt"

func main() {
	arr := []string{"a", "d", "f", "k"}
	target := "b"
	result := ceiling(arr, target)
	fmt.Println(result)

}

func ceiling(arr []string, target string) int {
	start := 0
	end := len(arr) - 1
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if target < arr[mid] {

			end = mid - 1
		} else if target > arr[mid] {

			start = mid + 1

		} else {

			return mid
		}
	}
	fmt.Println("start", start)
	fmt.Println("end", end)
	return start
}
