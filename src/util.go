package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
)

func testRun() bool {
	v := flag.Lookup("test.v")
	return v != nil
}

func printDebug(msg string) {
	fmt.Println("")
	fmt.Println("-------------------------------------")
	if msg != "" {
		fmt.Println(msg + "...")
	}
	fmt.Println("-------------------------------------")
}

func delayOutput() {
	outputDelay := os.Getenv("output_delay")
	if outputDelay != "" {
		if v, err := strconv.Atoi(outputDelay); err == nil {
			time.Sleep(time.Duration(v) * time.Millisecond)
		}
	} else {
		time.Sleep(1000 * time.Millisecond)
	}
}

func debugMsg(msg string) {
	if testRun() {
		return
	} else {
		printDebug(msg)
	}
}

func debugMsgDelay(msg string) {
	if testRun() {
		return
	} else {
		printDebug(msg)
		delayOutput()
	}
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
		debugMsgDelay("Creating a new deck")
		return newDeck()
	} else {
		debugMsgDelay("Loading a deck from file")
		return maybeDeck
	}
}
