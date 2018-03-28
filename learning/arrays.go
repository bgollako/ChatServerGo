package learning

import (
	"fmt"
	"reflect"
)

// MainArray contains examples of array
func MainArray() {
	fmt.Println("Hello Arrays")
	// Types of array declaration
	ar := [4]string{"Bhagirath", "Gollakota", "Sri", "Sruthi"}
	fmt.Println(ar)
	fmt.Println(reflect.TypeOf(ar))
	// Slicing an array also returns a slice
	sl := ar[1:3]
	fmt.Println(reflect.TypeOf(reflect.TypeOf(sl)))
	// In the below section the compiler will count the number for us
	array2 := [...]string{"Bhagavan", "Sathya", "Sai", "Baba"}
	fmt.Println(array2)
	// Get length of an array
	fmt.Println(len(array2))
	// Iterate through an array
	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}
	// Another way to iterate through an array
	for id, value := range array2 {
		fmt.Println(id, value)
	}

	// Use _ for non used variables as shown below
	for _, value := range array2 {
		fmt.Println(value)
	}
}
