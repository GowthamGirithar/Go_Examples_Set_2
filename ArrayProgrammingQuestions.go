package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var arr = [] int{1, 2, 4, 5, 7}
	FindMissingElementInArray()
	IsElementPresent(arr, 3)
	GuessingGame()
	FindDuplicateNumber()
	FindDuplicateCount()
	LargestAndSmallestNumber()
	PairIsEqualToGivenNumber()
	RemoveGivenElement(4)
	RemoveGivenElementWithoutAlteringOrder(4)
	RemoveDuplicatesFromArray()

}

func RemoveDuplicatesFromArray() {
	var arr = [] int{1, 4, 3, 2, 2}
	var duplicateCount = make(map[int]int) //1->1, 2-> 2, 3->1, 4->1
	for i := range arr {
		_, exists := duplicateCount[arr[i]]
		if exists {
			duplicateCount[arr[i]] += 1 //store index and count
		} else {
			duplicateCount[arr[i]] = 1
		}
	}
	var data []int
	for k, v := range duplicateCount {
		if v == 1 {
			data = append(data, k)
		}
	}
	fmt.Printf("%v", data)

}

func RemoveGivenElement(i int) {
	var arr = [] int{1, 4, 3, 2, 2}
	var numberIndex int
	for index := range arr {
		if arr[index] == i {
			numberIndex = index
		}
	}

	arr[numberIndex] = arr[len(arr)-1] // copy the last element to the number index
	arr[len(arr)-1] = 0                // invalidate the data
	arr = arr[:len(arr)-1]             // truncate
	fmt.Printf("%v", arr)
}
func RemoveGivenElementWithoutAlteringOrder(i int) {
	var arr = [] int{1, 4, 3, 2, 2}
	var numberIndex int
	for index := range arr {
		if arr[index] == i {
			numberIndex = index
		}
	}

	copy(arr[numberIndex:], arr[numberIndex+1:]) // [1,3,2,2,2]
	arr = arr[:len(arr)-1]                       // truncate
	fmt.Printf("%v", arr)
}

func PairIsEqualToGivenNumber() {
	var arr = [] int{1, 4, 3, 2, 2}
	var givenNumber int = 6
	var data = make([]int, len(arr))
	for i := range arr {
		isPresent := IsElementPresent(data, givenNumber-arr[i])
		if isPresent {
			fmt.Print("The pair is ", arr[i], givenNumber-arr[i])
			break
		}
		data = append(data, arr[i]) // check each time in slice and see whether the element is present after subtraccting

	}

}

// Smallest and largest number
func LargestAndSmallestNumber() {
	var arr = [] int{1, 2, 3, 2, 2}
	smallestNumber := math.MaxInt64 // math package have all the values
	largestNumber := 0
	for i := range arr {
		if arr[i] < smallestNumber {
			smallestNumber = arr[i]
		}
		if arr[i] > largestNumber {
			largestNumber = arr[i]
		}
	}
	println(" the smallest number and largest number is ", smallestNumber, largestNumber)
}

func FindDuplicateCount() {
	var arr = [] int{1, 2, 3, 2, 2}
	var duplicateCount = make(map[int]int)
	for i := range arr {
		_, exists := duplicateCount[arr[i]]
		if exists {
			duplicateCount[arr[i]] = duplicateCount[arr[i]] + 1
		} else {
			duplicateCount[arr[i]] = 1
		}
	}

	for k, v := range duplicateCount {
		if v > 1 {
			println("duplicated count is ", k, v)
		}
	}
}

func FindDuplicateNumber() {
	var arr = [] int{1, 2, 3, 5, 2} // cannot use [...]
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] == arr[i] {
			println("the duplicate number is ", arr[i])
		}
	}
}

//find the guessing number
func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

//element is present
func IsElementPresent(arr []int, num int) bool {
	// cannot use [...]
	value := sort.SearchInts(arr, num)
	//println(value) // print the index wheere it can be inserted , if present also it wll return the index
	if value < len(arr) && arr[value] == num {
		return true
	}
	return false
}

//missing number in array
func FindMissingElementInArray() {
	var arr = [] int{1, 2, 4, 5, 7} // cannot use [...]
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != 1 {
			println("the missing number is ", arr[i]+1)
		}
	}

}
