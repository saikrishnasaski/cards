package main

import "fmt"

func main() {

	cards := newDeck()

	cards.saveToFile("my_cards")

	cards = newDeckFromFile("my_cards")

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	fmt.Println()

	hand.shuffle()
	hand.print()
}

func newCard() string {

	return "Five of Diamonds"
}
