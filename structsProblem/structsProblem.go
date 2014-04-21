package main
import "fmt"

type Circle struct{
 r,x,y float32
}
func main() {
 var circle Circle;
 circle = Circle{r:10,x:12,y:35}
 fmt.Printf("Circle's Radius and Center(x,y) are::%f",circle.x)
}