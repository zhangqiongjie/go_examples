package main

func main() {
	cards := deck{"I am a student,", "9 years old", "class in 8", newCard()}
	cards = append(cards, "append haha")
	cards.print()
}

func newCard() string {
	return "keep good mood every miniute."
}
