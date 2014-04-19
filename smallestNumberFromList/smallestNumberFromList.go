package main
import "fmt"

func main() {
 x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 21,17,
	19,97, 59,17,
 }
 length := len(x)
 start := x[0]	//First Number of slice
 for i := 1; i < length; i++ {	//Loop through slice
  currentElem := x[i]	//Current Element in slice
  if start > currentElem {
   start = x[i]
  }
 }
 fmt.Println("Smallest Number Fro List ")
 fmt.Println(x);
 fmt.Println("   Is  :: ",start)
}