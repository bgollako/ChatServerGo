package learning

import (
	"fmt"
)

func mainSlices() {
	// Declaring slices
	slices := []string{"Bhagirath", "Gollakota"}
	fmt.Println(slices)
	nums := make([]int, 0, 0)
	fmt.Println("Capacity of nums is ", cap(nums))
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
		fmt.Println("Capacity of nums is ", cap(nums))
	}
	fmt.Println(nums)
	// slicing a slice
	num1to6 := nums[:6]
	fmt.Println(num1to6)
	// length - number of elements reffered to by the slices
	// Capacity - number of elements in the underlying array
	// slicing an array or slice creates a new slice that points to the original array.
	ns := num1to6[2:4]
	ns[0] = 100
	ns[1] = 1000
	fmt.Println(ns)
	fmt.Println(num1to6)
	fmt.Println(nums)
	// iterating through slices
	for index, value := range num1to6 {
		fmt.Println(index, value)
	}
}

func experimentWithSlices() {
	letters := make([]string, 5, 5)
	letters[0] = "a"
	letters[1] = "b"
	letters[2] = "c"
	letters[3] = "d"
	letters[4] = "e"
	// The below is panicing
	// letters[5] = "f"
	letters = append(letters, "f")
	letters = append(letters, "g")
	letters = append(letters, "h")
	letters = append(letters, "i")
	letters = append(letters, "j")
	letters = append(letters, "k")
	fmt.Println(letters)
}
