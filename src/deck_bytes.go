package main

import (
	"strings"
)

type deck_bytes []byte

func (db deck_bytes) toSlice() []string {
	byteSlice := []byte(db)
	deckString := string(byteSlice)
	return strings.Split(deckString, ",")
}
