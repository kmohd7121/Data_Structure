package main
import "fmt"
func main(){
	array:=[]int {10,20,30,40,50,60}
	k:=2
	fmt.Print("Original array before rotation\n")
	fmt.Print(array)
	z:=rotate(array,k)
	fmt.Print("\nafter rotation array\n")
	fmt.Print(z)
}
func rotate(data []int ,k int)[]int{
	n:=len(data)
	ReverseArray(data,0,k-1)
	ReverseArray(data,k,n-1)
	ReverseArray(data,0,n-1)
	return data
}
func 	ReverseArray(data []int,start int ,end int ) []int{
	i:=start
	j:=end
	for i<j{
		data[i],data[j]=data[j],data[i]
		i++
		j--
	}
	return data
}