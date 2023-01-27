package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	intList := nonRepeatingRandomArray(15)

	numberToSearch := rand.Intn(30)

	result := -1

	result = binarySearch(intList, numberToSearch)

	fmt.Printf("Searched number: %v\n", numberToSearch)
	fmt.Printf("Sorted list: %v\n", intList)

	if result == -1 {
		println("\nNumber not found")
	} else {
		fmt.Printf("\nNumber %v found at index #%v\n", numberToSearch, result)
	}
}

func binarySearch(intList []int, numToSearch int) int {
	left := 0
	right := len(intList) - 1

	for left <= right {
		mid := (left + right) / 2

		if intList[mid] == numToSearch {
			return mid
		}

		if intList[mid] < numToSearch {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func nonRepeatingRandomArray(length int) []int {
	// Create an empty array to hold the random numbers
	randomNumbers := make([]int, 0, length)

	// Fill the array with non-repeating random numbers
	for len(randomNumbers) < length {
		newNumber := rand.Intn(30)
		if !contains(randomNumbers, newNumber) {
			randomNumbers = append(randomNumbers, newNumber)
		}
	}

	// Sort the array before returning
	sort.Ints(randomNumbers)
	return randomNumbers
}

func contains(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
