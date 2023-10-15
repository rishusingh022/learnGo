package main

import "fmt"

func main() {
	cards := newCard()
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
