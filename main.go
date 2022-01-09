package main

import "fmt"

type person struct {
	first string
	last string
	favFlavors []string
}

func main() {
	p1 := person {
		first: "James",
		last: "Bond",
		favFlavors: []string{"Chcolate", "Vanilla"},
	}
	
	p2 := person {
		first: "Parathan",
		last: "Tl",
		favFlavors: []string{"Strawberry", "Milk Shake"},
	}
	
	m := map[string] person {
		p1.last: p1,
		p2.last: p2,
	}
	
	for _,v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("----------")
	}
	
}
