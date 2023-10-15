package main

import "fmt"

// var deckSize int

func main() {
	// var cards string = "Ace of Spades"
	// cards := "Ace of Spades"
	// cards = "Five of Diamonds"
	// cards := newCard()

	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	for i, cards := range cards {
		fmt.Println(i, cards)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
