package main

import "fmt"

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
}
