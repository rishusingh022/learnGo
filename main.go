package main

func main() {

	cards := newDeck()
	cards.print()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}
