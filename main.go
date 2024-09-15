package main

import "fmt"

func main() {

	var age int

	fmt.Println("enter age")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Eligible")
	} else {
		fmt.Println("Not eligible")
	}

}
