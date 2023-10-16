package main

func main() {

	// cards := newDeck()
	// cards.print()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
