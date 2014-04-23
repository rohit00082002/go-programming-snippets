package main
import "fmt"

func grfunc1( param int) {
	for i := 0; i< 10; i++ {
		fmt.Println("GRFunc-1::", i, " Parameter::", param)
	}
}

func grfunc2( param int) {
	for i := 0; i< 10; i++ {
		fmt.Println("GRFunc-2::", i, " Parameter::", param)
	}
}

func main() {
 go grfunc1(5)
 fmt.Println("Print After GORoutines1")
 go grfunc2(55)
 fmt.Println("Print After GORoutines2")
 var c string
 fmt.Scan(&c)
}