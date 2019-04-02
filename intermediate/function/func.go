package main

import "fmt"

// func max(i int, j int) int {
// 	if i > j {
// 		return i
// 	} else {
// 		return j
// 	}
// }

func max(i int, j int, k *int) {
	fmt.Println(k)
	if i > j {
		*k = i
	} else {
		*k = j
	}
}

func main() {
	var c int
	fmt.Println(&c)
	max(10, 15, &c)
	fmt.Println(c)
}
