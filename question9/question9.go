/* Question-9 == Prepare the code and flowchart that sorts the elements in a user-generated array from largest to smallest. */

package main

import "fmt"

func main() {
	arr := []int{29, 8, 5, 15, 32, 18, 1, 40, 63, 34}

	for i := 0; i < len(arr); i++ {

		for j := i; j < len(arr); j++ {

			if arr[j] < arr[i] {

				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)

}
