package main

import "fmt"

func main() {
	
	x := map[string][]string{
		"parathan": []string{"Jaffna", "SLIIT"},
		"YouTube": []string{"USA", "Google"},
		"Facebook": []string{"Ireland", "Mark"},
	}
	
	fmt.Println(x)
	
	x["Qrios"] = []string{"Golang", "React Native"}
	
	for k,v := range x {
		fmt.Println("This is the record for:", k)
		for i, v2 := range v {
			fmt.Println("\t",i, v2)
		}
	}
	
	fmt.Println("After Delete....")
	
	delete(x, "Qrios")
	
	for k,v := range x {
		fmt.Println("This is the record for:", k)
		for i, v2 := range v {
			fmt.Println("\t",i, v2)
		}
	}
}
