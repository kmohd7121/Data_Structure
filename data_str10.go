package main
import "fmt"
func main(){
	fmt.Print("enter a number \n")
	var x int 
	fmt.Scan(&x)
	z:=fib(x)
	fmt.Print(z)
}
func fib(n int)int {
	if n<=1{
		return n
	}
	return fib(n-1)+fib(n-2)
}