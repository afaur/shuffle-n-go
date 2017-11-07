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
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(hand deck, handSize int) (deck, deck) {
	newHand := append(hand, d[:handSize]...)
	newDeck := d[handSize:]
	return newDeck, newHand
}

func (d deck) toString() string {
	deckString := []string(d)
	return strings.Join(deckString, ",")
}

func (d deck) saveToFile(filename string) error {
	deckString := d.toString()
	byteSlice := []byte(deckString)
	return ioutil.WriteFile(filename, byteSlice, 0666)
}

func newDeckFromFile(filename string, softFail bool) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		if softFail != true {
			fmt.Println("Error:", err)
			os.Exit(1)
		} else {
			return deck{}
		}
	}
	deckBytes := deck_bytes(bs)
	deckSlice := deckBytes.toSlice()
	return deck(deckSlice)
}

func (d deck) shuffle() {
	slice := []string(d)
	Shuffle(slice)
	d = deck(slice)
}

func (d deck) dealCards(cardsPerPlayer int) (deck, deck, deck) {
	playerHand := deck{}
	dealerHand := deck{}

	cardsToDeal := 2 * cardsPerPlayer
	cardsDelt := 0

	fmt.Println("-------------------------------------")
	for cardsDelt < cardsToDeal {
		fmt.Println("Dealing one card to a player...")
		d, playerHand = d.deal(playerHand, 1)

		fmt.Println("Dealing one card to the dealer...")
		d, dealerHand = d.deal(dealerHand, 1)

		cardsDelt += 2
	}
	return d, playerHand, dealerHand
}
