package main

import "fmt"

func main() {

	x := [5]int{3,4,9,10,3}
	
	fmt.Println(x)
	
	for i, v := range x {
		fmt.Println(i,v)
	}
	
	fmt.Printf("%T",x)
}
