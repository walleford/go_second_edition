package main

import "fmt"

func main() {
	greetings := []string{
		"Hello",
		"Hola",
		"こんにちは",
		"Привіт",
	}

	greetingsSub := greetings[:2]
	greetingsSub2 := greetings[2:4]

	fmt.Println(greetings)
	fmt.Println(greetingsSub)
	fmt.Println(greetingsSub2)

	message := "hi"
	fmt.Println(message)

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	jordan := Employee{
		"Jordan",
		"Wallingsford",
		1,
	}
	suhayla := Employee{
		lastName:  "Sibaai",
		firstName: "Suhayla",
		id:        2,
	}
	var bata Employee
	bata.firstName = "Bata"

	fmt.Println(jordan)
	fmt.Println(suhayla)
	fmt.Println(bata)
}
