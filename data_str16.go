// Set implemented using map
package main
import "fmt"
type set map[interface {}]bool
func (s *set)Add(key interface{}){
	(*s)[key]=true
}
func (s *set)Remove(key interface{}){
	delete((*s),key)
}
func (s *set)find(key interface{})bool{
	return (*s)[key]
}
func main(){
	mp:=make(set)
	mp.Add("a")
	mp.Add("b")
	mp.Add("a")
	fmt.Println(mp.find("a"))
	fmt.Println(mp.find("b"))
	fmt.Println(mp.find("c"))
}