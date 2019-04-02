package main

import (
	"fmt"
)

func main() {
	age := 55
	if age < 13 {
		fmt.Println("Your are young")
	} else if age < 20 {
		fmt.Println("Your are teen")
	} else if age < 30 {
		fmt.Println("Your are twenties")
	} else if age < 40 {
		fmt.Println("Your are therties")
	} else if age < 50 {
		fmt.Println("Your are fourthies")
	} else {
		fmt.Println("Your are old")
	}
}
