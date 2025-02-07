package main

import (
	"fmt"
	// "log"
)

func main() {
	// var maxnum int
	maxnum := 25

	// fmt.Println("number of integers: ")
	// _, err := fmt.Scanf("%d", &maxnum)
	// if err != nil {
	// log.Fatal(err)
	// }

	for i := 1; i < (maxnum + 1); i++ {
		// fmt.Println(i)
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			if i%3 == 0 {
				fmt.Println("fizz")
			} else {
				if i%3 == 5 {
					fmt.Println("buzz")
				} else {
					fmt.Println(i)
				}
			}

		}

	}

}
