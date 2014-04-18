package main	//Main Package
import "fmt" 	//importing "Format" package

func main() {
 var method string	//String variable
 var a ,b float32	//Float variable
 fmt.Print("Enter Method as '+', '-', '*', '/'   ")	//Output on screen
 fmt.Scan(&method)	//User input
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