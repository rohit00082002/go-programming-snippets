package main
import "fmt"

func main() {
 counter := 0
 for i:= 1; i <= 100; i++ {
  if i % 3 == 0 {
   counter++
   fmt.Print(i,"     ")

   if counter % 10 == 0 {
    fmt.Println("")
	fmt.Println("")
   }

   }
  
 }
}