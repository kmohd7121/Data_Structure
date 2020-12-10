// Binary search program 
// This program run  only for sort array
package main

import "fmt"

func main() {
	fmt.Print("enter how many number you want in array\n")
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Print("enter the element of array\n")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Print("enter a number you want to search in array\n")
	var x int
	fmt.Scan(&x)
	z := BinarySearch(arr, n, x)
	fmt.Print(z)
}
func BinarySearch(data []int, size int, value int) bool {
	var mid int
	low := 0
	heigh := size - 1
	for low <= heigh {
		mid = low + (heigh-low)/2
		if data[mid] == value {
			return true
		} else {
			if data[mid] < value {
				low = mid + 1
			} else {
				heigh = mid - 1
			}
		}
	}
	return false
}
