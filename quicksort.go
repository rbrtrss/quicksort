package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	quicksort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}

func makeRandomSlice(numItems, max int) []int {
	s := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		s[i] = rand.Intn(max)
	}
	return s
}

func printSlice(slice []int, numItems int) {
	if numItems > len(slice) {
		numItems = len(slice)
	}
	s := slice[:numItems]
	fmt.Printf("%d \n", s)
}

func checkSorted(slice []int) {
	sorted := true
	msg := "The slice is sorted"
	for i := 0; i < (len(slice) - 1); i++ {
		if slice[i] > slice[i+1] {
			sorted = false
			break
		}
	}
	if !sorted {
		msg = "The slice is NOT sorted!"
	}
	fmt.Println(msg)
}

func partition(slice []int) int {
	// lo := 0 always
	i := -1 // lo - 1
	hi := len(slice) - 1
	pivot := slice[hi]

	for j := 0; j < hi; j++ {
		jValue := slice[j]
		if jValue <= pivot {
			i++
			iValue := slice[i]
			// swaping
			slice[j] = iValue
			slice[i] = jValue
		}
	}
	i++
	// swaping
	iValue := slice[i]
	jValue := slice[hi]
	slice[hi] = iValue
	slice[i] = jValue
	return i
}

func quicksort(slice []int) {
	if len(slice) > 1 {
		p := partition(slice)

		quicksort(slice[:p])
		quicksort(slice[p+1:])
	}
}
