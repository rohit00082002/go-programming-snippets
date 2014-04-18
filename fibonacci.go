package main
import "fmt"
func main() {
 var a,b = 0,1
 fmt.Printf(" %d ", a)
 for i := 1; i<10; i++ {
 fmt.Printf(" %d ", b)
  b,a = b + a, b;
 }
}