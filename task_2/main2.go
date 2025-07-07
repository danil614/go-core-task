package task_2

import (
	"fmt"
	"math/rand"
)

func getRandomNumSlice() []int {
	originalSlice := make([]int, 10)
	for i := 0; i < len(originalSlice); i++ {
		originalSlice[i] = rand.Intn(1000) - 500
	}
	return originalSlice
}

func sliceExample(slice []int) []int {
	if slice == nil {
		return nil
	}

	var evenSlice []int
	for _, v := range slice {
		if v%2 == 0 {
			evenSlice = append(evenSlice, v)
		}
	}
	return evenSlice
}

func addElements(slice []int, num int) []int {
	if slice == nil {
		return nil
	}
	return append(slice, num)
}

func copySlice(slice []int) []int {
	if slice == nil {
		return nil
	}
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func removeElement(slice []int, index int) []int {
	if slice == nil {
		return nil
	}
	if index < 0 || index >= len(slice) {
		panic("index out of range")
	}

	return append(slice[:index], slice[index+1:]...)
}

func Main() {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Random slice:")
	fmt.Println(getRandomNumSlice())

	fmt.Println("\nEven slice:")
	fmt.Println(sliceExample(originalSlice))

	fmt.Println("\nAdd element:")
	fmt.Println(addElements(originalSlice, 11))

	fmt.Println("\nCopy slice:")
	copySlice := copySlice(originalSlice)
	fmt.Println(copySlice)
	originalSlice[0] = 11
	fmt.Println("After append:", copySlice)
	fmt.Println("Original slice:", originalSlice)

	fmt.Println("\nDelete element (i=2):")
	fmt.Println(removeElement(originalSlice, 2))
}
