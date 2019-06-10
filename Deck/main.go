package main

func main() {
	cards := newDeck()

	cards.shuffle()
	cards.print()
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("myFile")

	// cards := newDeckFromFile("myFile")
	// cards.print()
}
