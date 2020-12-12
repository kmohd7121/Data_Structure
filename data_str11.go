//Search a value in an increasing order sorted list of integers.
package main
import "fmt"
func main(){
	arr:=[]int {1,2,3,4,5,6,7,8}
	size:=len(arr)
	low:=0
	heigh:=size-1
	fmt.Print("enter anumber you want to search in array\n")
	var x int
	fmt.Scan(&x)
	fmt.Print(rec(arr,low,heigh,x))

}
func rec(arr []int ,low int ,high int ,value int)int{
	mid:=low+(high-low)/2
	if arr[mid]==value{
		return mid
	}else if arr[mid]<value{
		return rec(arr, mid+1, high, value)
	}else{
		return rec(arr, low, mid-1, value)
	}
	
}
