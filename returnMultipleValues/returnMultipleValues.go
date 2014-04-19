package main
import "fmt"

func returnMultipleValues() (int, int) {
	i := 3
	j := 5
	return i, j
}

func main() {
 x,y := returnMultipleValues()
 fmt.Println(x);
 fmt.Println(y);
}