package main
import "fmt"

func main() {
 var method string
 var a ,b float32
 fmt.Print("Enter Method as '+', '-', '*', '/'   ")
 fmt.Scan(&method)
 fmt.Print("Enter First Number   ")
 fmt.Scan(&a)
 fmt.Print("Enter Second Number   ")
 fmt.Scan(&b)
 switch method {
  case "+": 
	fmt.Printf("%f + %f = %f ", a, b, a + b )
	break
  case "-":
	fmt.Printf("%f - %f =  %f ", a, b, a - b )
	break
  case "*":
	fmt.Printf("%f * %f =  %f ", a, b, a * b )
	break
  case "/":
	fmt.Printf("%f / %f =  %f ", a, b, a / b )
	break
  default: fmt.Print("Invalid Method")
	break
 }
}