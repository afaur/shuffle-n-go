package main

func main() {
	debugMsg("Starting deck simulation")
	deck := createOrLoadDeck("decks/new_deck")

	debugMsg("Attempting to Shuffle deck")
	deck.shuffle()

	debugMsg("Saving the shuffled deck")
	deck.saveToFile("decks/shuffled_deck")

	debugMsg("Dealing 7 cards to each player")
	deck, playerHand, dealerHand := deck.dealCards(7)

	debugMsg("Cards in virtual players hand")
	playerHand.print()

	debugMsg("Cards currently in dealers hand")
	dealerHand.print()

	debugMsg("Cards left in deck")
	deck.print()
}
