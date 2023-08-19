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

	fmt.Println(hand.toString())
}

func newCard() string {

	return "Five of Diamonds"
}
