package main

import "fmt"

func main() {
	for i := 1; i < 21; i++ {
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

	// if i%3 == 0 {
	// fmt.Println("fizz")
	// } else {
	// fmt.Println(i)
	// }
	// else {
	// if i%5 == 0 {
	// fmt.Println("buzz")
	// }
	// }
}
