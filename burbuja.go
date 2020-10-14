package main

import "fmt"

// Swap Function
func swapArray(a *int64, b *int64) {
	var temp int64
	temp = *a // Dereference
	*a = *b
	*b = temp
}

// Bubble Function
func Burbuja(s []int64) {
	for i := len(s) - 1;  i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[j] > s[j + 1] {
				swapArray(&s[j], &s[j + 1]) // Direction
			}
		}
	}
}

// Main Function
func main()  {
	// Array
	array := []int64{1, 3, 5, 6, 7, 8, 9, 0, 9, -1, -5}
	// Original
	fmt.Print(array)
	// Sort
	Burbuja(array)
	// result
	fmt.Print(array)
}
