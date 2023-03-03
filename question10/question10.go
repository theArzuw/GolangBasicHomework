/* Question-10 == Given an array, prepare a code and flowchart that gives the number of elements in the array.
For example :
Input : arr[] = {3, 3, 3, 3, 2, 3, 4, 1}
Output : 3 "3" elements, 1 "2" element, 1 "4" element, 1 "1" element. */

package main

import "fmt"

func main() {
	arr := []int{30, 60, 40, 70, 30, 50, 10, 80, 40, 50, 70, 20}
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num] = freq[num] + 1
	}
	fmt.Println("Frequency of the Array is : ", freq)
}
