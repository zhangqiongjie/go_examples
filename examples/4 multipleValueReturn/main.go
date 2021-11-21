package main

func main() {
	cards := newCard()
	hand, remainingCard := deal(cards, 5)
	hand.print()
	remainingCard.print()
}
