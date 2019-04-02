package main

import "fmt"

// const maxCap int = 5

// var listGroceries [maxCap]string
var listGroceries []string

// var curCap int

// func addGrocery(a string) {
// 	fmt.Println("Capacity", cap(listGroceries))
// 	listGroceries = append(listGroceries, a)
// 	// if curCap < maxCap {
// 	// 	listGroceries[curCap] = a
// 	// 	curCap++
// 	// } else {
// 	// 	fmt.Println("Full")
// 	// }
// }

func addGrocery(a ...string) {
	fmt.Println("Capacity", cap(listGroceries))
	for _, d := range a {
		listGroceries = append(listGroceries, d)
	}
	// if curCap < maxCap {
	// 	listGroceries[curCap] = a
	// 	curCap++
	// } else {
	// 	fmt.Println("Full")
	// }
}

func basketGrocery() {
	fmt.Println("List Grocery:")
	// for i := 0; i < len(listGroceries); i++ {
	// 	fmt.Println(listGroceries[i])
	// }
	// for i, d := range listGroceries {
	// 	fmt.Println(i, d)
	// }
	for _, d := range listGroceries {
		fmt.Println(d)
	}
}

func main() {
	addGrocery("Bread")
	addGrocery("Cucumber")
	addGrocery("Cake")
	addGrocery("Apple")
	addGrocery("Fruit Cake")
	addGrocery("Pokemon Game")
	addGrocery("Ice Cream", "Beans")
	basketGrocery()
}
