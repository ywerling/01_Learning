package main

import (
	"fmt"
	// "log"
	// "log"
)

func main() {
	// var maxnum int
	maxnum := 25

	// var maxnum int
	//
	// fmt.Print("Enter an integer value: ")
	// _, err := fmt.Scanf("%d", &maxnum)
	//
	// if err != nil {
	// log.Fatal("Invalid input! Please enter an integer.") // Terminate on error
	// }
	//
	fmt.Println("You have entered:", maxnum)

	for i := 1; i < (maxnum + 1); i++ {
		// fmt.Println(i)
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			if i%3 == 0 {
				fmt.Println("fizz")
			} else {
				if i%5 == 0 {
					fmt.Println("buzz")
				} else {
					fmt.Println(i)
				}
			}

		}

	}

}
