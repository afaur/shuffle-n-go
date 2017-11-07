package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	cardCount := 52
	if len(deck) != cardCount {
		errMsgExp := "Expected new deck to have %v card(s)"
		errMsgAct := "Actual new deck has %v card(s)"
		t.Errorf(errMsgExp, cardCount)
		t.Errorf(errMsgAct, len(deck))
	}
}

func TestNewDeckEachCard(t *testing.T) {
	actualDeck := newDeck()
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	idx := 0
	for _, suit := range suits {
		for _, value := range values {
			expCard := value + " of " + suit
			actCard := actualDeck[idx]
			if expCard != actCard {
				errMsgExp := "Expected card @ pos %v to be %v"
				errMsgAct := "Actual card @ pos %v was %v"
				t.Errorf(errMsgExp, idx, expCard)
				t.Errorf(errMsgAct, idx, actCard)
			}
			idx++
		}
	}
}

func TestDealCards(t *testing.T) {
	actualDeck := newDeck()
	playerHand := deck{}

	cardsToDeal := 5

	deltDeck, playerHand := actualDeck.deal(playerHand, cardsToDeal)

	if len(playerHand) != cardsToDeal {
		errMsgExp := "Expected player hand to have %v card(s)"
		errMsgAct := "Actual player hand has %v card(s)"
		t.Errorf(errMsgExp, cardsToDeal)
		t.Errorf(errMsgAct, len(playerHand))
	}

	deltDeckLenAct := len(deltDeck)
	deltDeckLenExp := len(actualDeck) - cardsToDeal

	if deltDeckLenAct != deltDeckLenExp {
		errMsgExp := "Expected delt deck to have %v card(s)"
		errMsgAct := "Actual delt deck has %v card(s)"
		t.Errorf(errMsgExp, deltDeckLenExp)
		t.Errorf(errMsgAct, deltDeckLenAct)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFile, cardCount := "._testDeck", 52
	os.Remove(testFile)
	deck := newDeck()
	deck.saveToFile(testFile)
	loadedDeck := newDeckFromFile(testFile, false)
	if len(loadedDeck) != cardCount {
		errMsgExp := "Expected loaded deck to have %v card(s)"
		errMsgAct := "Actual loaded deck has %v card(s)"
		t.Errorf(errMsgExp, cardCount)
		t.Errorf(errMsgAct, len(loadedDeck))
	}
	os.Remove(testFile)
}
