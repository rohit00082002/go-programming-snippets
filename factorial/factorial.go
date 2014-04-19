package main
import "fmt"

func factorial(num int) int {
 if num < 1 {
	return 1
 } else {
  return num * factorial( num -1)
 }
}

func main() {
 var inputNo int

 fmt.Println("Please enter number to find its factorial")
 fmt.Scan(&inputNo)
 fmt.Println( factorial(inputNo) )
}