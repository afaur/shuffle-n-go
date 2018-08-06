package main

import "os"

func main() {
	os.Setenv("output_delay", "500")
	cardsPerPlayer := 5

	debugMsgDelay("Starting deck simulation")
	deck := createOrLoadDeck("decks/new_deck")

	debugMsgDelay("Attempting to Shuffle deck")
	deck.shuffle()

	debugMsgDelay("Saving the shuffled deck")
	deck.saveToFile("decks/shuffled_deck")

	debugMsgDelay("Dealing 7 cards to each player")
	playerHand, dealerHand := deck.dealCards(cardsPerPlayer)

	debugMsgDelay("Cards in virtual players hand")
	playerHand.print()

	debugMsgDelay("Cards currently in dealers hand")
	dealerHand.print()

	debugMsgDelay("Cards left in deck")
	deck.print()
}
