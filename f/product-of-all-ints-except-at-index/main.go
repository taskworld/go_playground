package main

import "fmt"

func main() {
	input := []int{3, 1, 2, 5, 6, 4}

	output := productOfAllIntsExceptAtIndex(input)

	for _, item := range output {
		fmt.Print(item, " ")
	}

	fmt.Println("")

	output = productOfAllIntsExceptAtIndexWithNoDivision(input)

	for _, item := range output {
		fmt.Print(item, " ")
	}
}

func productOfAllIntsExceptAtIndex(arr []int) (result []int) {

	result = make([]int, len(arr))

	product := 1
	for _, item := range arr {
		product *= item
	}

	for index, item := range arr {
		result[index] = product / item
	}

	return
}

func productOfAllIntsExceptAtIndexWithNoDivision(arr []int) (result []int) {

	result = make([]int, len(arr))

	product := 1
	for index, item := range arr {
		result[index] = product
		product *= item
	}

	product = 1
	for index := range arr {
		reverseIndex := len(arr) - 1 - index
		result[reverseIndex] = result[reverseIndex] * product
		product *= arr[reverseIndex]
	}

	return result
}
