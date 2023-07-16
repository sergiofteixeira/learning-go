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

//same as above but inline
//func main() {
//for i := 1; i <= 10; i++ {
//fmt.Println(i)
//}
//}

// for each together with a condition
//func main() {
	//i := 1
	//for i <= 10 {
		//if i%2 == 0 {
			//fmt.Println(i, "even")
		//} else {
            //fmt.Println(i, "odd")
        //}
		//i += 1
	//}
//}

//switch
func main() {
	i := 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
            fmt.Println(i, "odd")
        }
		i += 1
	}
}
