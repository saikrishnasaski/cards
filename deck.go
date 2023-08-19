package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {

		for _, value := range cardValues {

			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) print() {

	for i, card := range d {

		fmt.Println(i, card)
	}

}

func (d deck) toString() string {

	cards := []string(d)

	return strings.Join(cards, ",")
}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {

		fmt.Println("Error: ", err)

		os.Exit(1)
	}

	cards := strings.Split(string(bytes), ",")

	return deck(cards)

}
