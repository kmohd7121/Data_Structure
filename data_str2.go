//Write a method, which will search a list for some given value.
package main
import "fmt"

func SequentialSearch(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

func main(){
	arr:=[]int {11,12,13,34,23,24,45}
	fmt.Print("enter a number to search in array\n")
	var x int
	fmt.Scan(&x)
	z:=SequentialSearch(arr,x)
	fmt.Print(z)

}