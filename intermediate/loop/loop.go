package main

import (
	"fmt"
	"time"
)

func main() {

	// while loop
	i := 10
	for i > 0 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i--
	}

	//for loop
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("Happy new Year")
}
