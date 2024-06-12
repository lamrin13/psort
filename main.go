package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func generateSlice(lenght int16) []int32 {
	array := make([]int32, lenght)
	for i := range array {
		array[i] = rand.Int31()
	}
	return array
}

func findLargestNum(array []int32) int32 {
	largestNum := int32(0)
	for i := 0; i < len(array); i++ {
		if array[i] > largestNum {
			largestNum = array[i]
		}
	}
	return largestNum
}

func radixSort(array []int32) []int32 {
	largestNum := findLargestNum(array)
	size := len(array)
	significantDigit := int32(1)
	semiSorted := make([]int32, size, size)

	for largestNum/significantDigit > 0 {
		bucket := [10]int{0}
		for i := 0; i < size; i++ {
			bucket[(array[i]/significantDigit)%10]++
		}
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}
		for i := size - 1; i >= 0; i-- {
			bucket[(array[i]/significantDigit)%10]--
			semiSorted[bucket[(array[i]/significantDigit)%10]] = array[i]
		}
		for i := 0; i < size; i++ {
			array[i] = semiSorted[i]
		}
		significantDigit *= 10
	}

	return array
}

func main() {
	size := int16(10)
	input := generateSlice(size)
	input1 := make([]int32, size)
	copy(input1, input)
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	fmt.Println(input)
	fmt.Println(radixSort(input1))
}
