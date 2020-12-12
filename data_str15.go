package main
import "fmt"
import "github.com/golang-collections/collections/set"

func main() {
	st := set.New()
	st.Insert(1)
	st.Insert(2)
	fmt.Println(st.Has(1))
	fmt.Println(st.Has(3))
}
