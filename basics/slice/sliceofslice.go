package main

import "fmt"

func main() {
	
	x := []string{"James", "Bond", "Shaken, not stimmed"}
	y := []string{"Miss", "Moneypenny", "Hello James"}
	z := [][]string{x,y}
	
	fmt.Println(z)
	
	for i, xs := range z {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: \t %v \t value: \t %v \n", j, val)
		}
	}
	
}
