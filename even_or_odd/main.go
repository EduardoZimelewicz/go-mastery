package main

import "fmt"

func main() {

	slice := generateSlice()

	for index, _ := range slice {
		if checkEven(slice[index]) {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

}

func checkEven(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}

func generateSlice() []int {
	slice := []int{}

	for i := 0; i <= 10; i++ {
		slice = append(slice, i)
	}

	return slice
}
