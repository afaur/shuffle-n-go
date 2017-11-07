package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func debugMsg(msg string) {
	fmt.Println("")
	fmt.Println("-------------------------------------")
	if msg != "" {
		fmt.Println(msg + "...")
	}
	fmt.Println("-------------------------------------")
	time.Sleep(2500 * time.Millisecond)
}

func Shuffle(s interface{}) {
	rand.Seed(time.Now().UnixNano())
	value, swapPositions := GetValueAndSwapper(s)
	lIndex := value.Len() - 1
	for cIndex := lIndex; cIndex > 0; cIndex-- {
		rPosition := RandomInt(0, cIndex+1)
		swapPositions(cIndex, rPosition)
	}
}

func GetValueAndSwapper(s interface{}) (reflect.Value, func(idx, rpl int)) {
	return reflect.ValueOf(s), reflect.Swapper(s)
}

func RandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

func createOrLoadDeck(deckFile string) deck {
	maybeDeck := newDeckFromFile(deckFile, true)
	if len(maybeDeck) == 0 {
		debugMsg("Creating a new deck")
		return newDeck()
	} else {
		debugMsg("Loading a deck from file")
		return maybeDeck
	}
}
