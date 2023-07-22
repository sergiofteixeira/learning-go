package main

import (
	"fmt"
)

// basic
//func main() {
//i := 1
//for i <= 10 {
//fmt.Println(i)
////i = i + 1
//i += 1
//}
//}

// same as above but inline
func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
