package main

import "fmt"

func main() {
	for day := 1; day <= 12; day++ {
		fmt.Println("On the ", day, " Of Chrismat my true love sent to me:")
		switch day {
		case 12:
			fmt.Println("12 Drummers Drumming")
			fallthrough
		case 11:
			fmt.Println("11 Drummers Drumming")
			fallthrough
		case 10:
			fmt.Println("10 Drummers Drumming")
			fallthrough
		case 9:
			fmt.Println("9 Drummers Drumming")
			fallthrough
		case 8:
			fmt.Println("8 Drummers Drumming")
			fallthrough
		case 7:
			fmt.Println("7 Drummers Drumming")
			fallthrough
		case 6:
			fmt.Println("6 Drummers Drumming")
			fallthrough
		case 5:
			fmt.Println("5 Drummers Drumming")
			fallthrough
		case 4:
			fmt.Println("4 Drummers Drumming")
			fallthrough
		case 3:
			fmt.Println("3 Drummers Drumming")
			fallthrough
		case 2:
			fmt.Println("2 Drummers Drumming")
			fallthrough
		case 1:
			fmt.Println("1 Drummers Drumming")
		default:
			fmt.Println("0 Drummers Drumming")
		}
	}
}
