// Binary search in a sorted list.

package main

import "fmt"

func main(){
	arr:=[]int {1,2,3,4,5,6,7,8,9}
	fmt.Print("enter a number you want to search in array\n")
	var x int
	fmt.Scan(&x)
	z:=BinarySearch(arr,x)
	fmt.Print(z)

	
}

func BinarySearch(data []int,value int)bool{
	size :=len(data)
	var mid int
	low:=0
	high:=size-1
	for low<=high{
		mid=low+(high-low)/2
		if data[mid]==value{
			return true
		}else{
			if data[mid]<value{
				low=mid+1
			}else{
				high=mid-1
			}
		}
	}
	return false
		
}