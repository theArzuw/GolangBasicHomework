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
