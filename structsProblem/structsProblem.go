package main
import (
	"fmt"; "math"
)
/*
type Shape interface {
 area() float64
}
*/

// Circle Struct has scope outside of package, because it starts with Capital Letter.
type Circle struct{
 r,x,y float64
}

type Rectangle struct {
 w,h float64
}

func ( rectangle *Rectangle) area() float64 {
 return rectangle.w * rectangle.h
}

func ( circle *Circle) area() float64 {
 return math.Pi * circle.r * circle.r
}



func main() {
 var circle Circle
 var rectangle Rectangle

 circle = Circle{r:7,x:12,y:35}
 rectangle = Rectangle{w:10, h:25}

 circleArea := circle.area()
 fmt.Println("AREA   ", circleArea)

 rectangleArea := rectangle.area()
 fmt.Println("Rectangle area   ", rectangleArea)

}