package main
import (
	"fmt"
)
func main() {
	fmt.Println("Hello, Black Hat Gophers!")
	var count = int(42)
	ptr := &count
	fmt.Println(*ptr)
	fmt.Println(ptr)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		
	}
}