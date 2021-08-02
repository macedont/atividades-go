package main

import "fmt"

type vasco int

var x vasco

func algo(x interface{}) {
	fmt.Printf("%v %T", x, x)
}

func main() {
	x = 12
	algo(x)
}
