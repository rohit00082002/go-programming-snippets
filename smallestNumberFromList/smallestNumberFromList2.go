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
 start := x[0]	//First Number of slice
 for _ , v := range x {	//Loop through slice
  currentElem := v	//Current Element in slice
  if start > currentElem {
   start = v
  }
 }
 fmt.Println("Smallest Number Fro List ")
 fmt.Println(x);
 fmt.Println("   Is  :: ",start)
}