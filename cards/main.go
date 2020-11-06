package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
	cards.saveToFile("my_cards")
	hand, remainingDeck := deal(cards, 3)

	fmt.Println("Cards in hand============>")
	hand.print()

	fmt.Println("Remainig cards============>")
	remainingDeck.print()

}
