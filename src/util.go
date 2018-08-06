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

func printFancyDebug(msg string) {
	fmt.Println("-------------------------------------")
	if msg != "" {
		fmt.Println(msg + "...")
	}
	fmt.Println("-------------------------------------")
	fmt.Println("")
}

func printDebug(msg string) {
	if msg != "" {
		fmt.Printf(msg)
	}
}

func delayAmountInMs() int {
	delayMs := 1000
	outputDelay := os.Getenv("output_delay")
	if v, err := strconv.Atoi(outputDelay); err == nil {
		delayMs = v
	}
	return delayMs
}

func delayOutput() {
	time.Sleep(time.Duration(delayAmountInMs()) * time.Millisecond)
}

func debugMsgCall(msg string) {
	if msg == "" {
		println("\n")
	} else {
		if msg[:3] == "CLS" {
			delayOutput()
			fmt.Printf("\x1b[2K\r")
		} else if msg[:1] == "\b" {
			printDebug(msg)
		} else {
			printFancyDebug(msg)
		}
	}
}

func debugMsg(msg string) {
	if testRun() {
		return
	} else {
		debugMsgCall(msg)
	}
}

func debugMsgDelay(msg string) {
	if testRun() {
		return
	} else {
		debugMsgCall(msg)
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
