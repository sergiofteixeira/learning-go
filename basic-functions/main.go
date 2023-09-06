package main

import "fmt"

//func main() {
//fmt.Print("Enter temperature in fahrenheit: ")
//var (
//input float64
//)
//fmt.Scanf("%f", &input)
//output := (input - 32) / 1.8
//fmt.Println(output)
//}

func main() {
	fmt.Print("Enter number of feets: ")
	var (
		input float64
	)
	fmt.Scanf("%f", &input)
	output := input * 0.3048
	fmt.Println(output, " meters")
}
