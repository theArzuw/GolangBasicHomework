/* Question-3 == Prepare the code and flowchart that prints the sum of the numbers from 1 to 100 on the screen. */
package main

import "fmt"

func main() {
	//Calculate the sum of wings from 1 to 100

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	//Print the sum to the screen
	fmt.Println("1 ile 100 arasındaki sayıların toplamı: ", sum)
}
