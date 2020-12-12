package main
import "fmt"
func main(){
	fmt.Print("Greatest common divisor number is\n")
	z:=GCD(6,12)
	fmt.Print("Greatest common divisor ",z)


}
func GCD(m int, n int) int {
	if m < n {
		return GCD(n, m)
	}
	if m%n == 0 {
	    return n
	}
	return GCD(n,m%n)
}