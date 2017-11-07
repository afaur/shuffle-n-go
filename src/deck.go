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

func (d *deck) deal(hand *deck, handSize int) {
	curDeck := *d
	newHand := *hand
	*hand = append(newHand, curDeck[:handSize]...)
	*d = curDeck[handSize:]
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

func (d *deck) dealCards(cardsPerPlayer int) (deck, deck) {
	playerHand := deck{}
	dealerHand := deck{}

	cardsToDeal := 2 * cardsPerPlayer
	cardsDelt := 0

	for cardsDelt < cardsToDeal {
		debugMsg("Dealing one card to the player")
		d.deal(&playerHand, 1)

		debugMsg("Dealing one card to the dealer")
		d.deal(&dealerHand, 1)

		cardsDelt += 2
	}
	return playerHand, dealerHand
}
