// map/dictionary implementation in go collection

package main
import "fmt"
func main(){
	m:=make(map[string]int)
	m["Apple"]=40
	m["Banana"]=30
	m["mango"]=50
	for key,value:=range m{
		fmt.Print("[",key,"->",value,"]\n")
	}
	fmt.Println("Apple Price:",m["Apple"])
	delete(m,"Apple")
	fmt.Println("Apple Price:",m["Apple"])
	v,ok:=m["Apple"]
	fmt.Println("Apple price:",v,"Present:",ok)
	v2,ok2:=m["Banana"]
	fmt.Println("Banana price:",v2,"Present:",ok2)
}