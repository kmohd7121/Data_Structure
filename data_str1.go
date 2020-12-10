//Write a method that will return the sum of all the elements of the integer list,
given list as an input argument.
package main
import "fmt"
func main(){
	arr:=[]int {1,23,4,54,6}
	fmt.Print("sum of array is  ",sum(arr))
}
func sum(arr []int)int{
	size:=len(arr)
	total:=0
	for i:=0;i<size;i++{
		total+=arr[i]
	}
	return total
}