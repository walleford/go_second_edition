package main

import "fmt"

func main() {
	intSlice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		intSlice = append(intSlice, i)
	}

	for _, v := range intSlice {
		if v%3 == 0 && v%2 == 0 {
			fmt.Println("six")
			break
		}
		if v%2 == 0 {
			fmt.Println("two")
		}
		if v%3 == 0 {
			fmt.Println("three")
		}
	}

}

func shadowing() {
	x := 10 // shadowed var
	if x > 5 {
		fmt.Println(x) // prints 10
		x := 5         // creates a shadowing variable
		fmt.Println(x) // prints 5
	}
	fmt.Println(x) // prints 10
}

func shadowing_multiple_assignment() {
	x := 10 // shadowed
	if x > 5 {
		x, y := 5, 20     // shadows x in outer block
		fmt.Println(x, y) // prints 5 20
	}
	fmt.Println(x) // prints 10
}

// # switch-please
func switch_please() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length: ", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word...")
		}
	}
}
