package main

import "fmt"

func main() {

	//Get the student visa and final grade
	var visa, final int
	fmt.Println("Please enter visa note: ")
	fmt.Scanln(&visa)
	fmt.Println("Please enter final grade: ")
	fmt.Scanln(&final)

	//Calculate 30% of the visa grade

	visaContribution := float64(visa) * 0.3

	// Calculate 70% of final grade
	finalContribution := float64(final) * 0.7

	//Collect 30% of the visa grade and 70% of the final grade
	average := visaContribution + finalContribution/2

	//Print average grade to screen
	fmt.Println("Average: ", average)

	//Ask if you pwant to do another calculation
	fmt.Println("Do you want to create another account? (Yes/No)")
	var answer string
	fmt.Scanln(&answer)

	//Back to top if user wants to do another calculation
	if answer == "Yes" {

		//Finish if the user doesn't want to do any other calculations

	} else if answer == "No" {
		fmt.Println("Process completed. We thank you.")

	}
}
