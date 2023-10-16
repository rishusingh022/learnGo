package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// cards := newDeck()
	// cards.print()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// Assignment solution
	arrayNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, num := range arrayNumber {
		if num%2 == 0 {
			fmt.Printf("Number %v is even\n", num)
		} else {
			fmt.Printf("Number %v is odd\n", num)
		}
	}

	//1st method
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	//2nd method
	alex1 := person{"Alex", "Anderson", contactInfo{}}
	fmt.Println(alex1)

	//3rd method
	var alex3 person
	// fmt.Println(alex3)

	alex3.firstName = "Alex"
	alex3.lastName = "Anderson"
	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)
}
