package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func main() {
	//one()
	//two()
	//three()
	//four()
	//five()
	//six()
	//seven()
	//eight()
	//nine()
	ten()
}

func one() {
	loopfor()
}
func two() {
	customTypeDeck()
}
func three() {
	cards := newDeck()
	cards.print()
}
func four() {
	cards := newDeck()
	hand, remainingCard := deal(cards, 10)
	hand.print()
	fmt.Println("-----")
	remainingCard.print()
}
func five() {
	cards := newDeck()
	fmt.Println(cards.toString())
}
func six() {
	cards := newDeck()
	cards.saveToFile("my-cards.txt")
}
func seven() {
	cards := readFromFile("my-cards.txt")
	cards.print()
}
func eight() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
func nine() {
	getMap()
}
func ten() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(sb)
	printGreeting(eb)
}

//1
func loopfor() {
	cards := []string{"I am learning for loop,", printState()}
	cards = append(cards, "I am value of append array.")
	fmt.Print(cards)
	for s, card := range cards {
		fmt.Println(s, card)
		//fmt.Print(s, card)
	}
}

func printState() string {
	return "successed."
}

//2
func customTypeDeck() {
	cards := deck{"I am s student,", "9 years old,", "class in 8.", newCard()}
	cards = append(cards, "I am append value.")
	cards.print()
}

func newCard() string {
	return "keep good mood every miniute."
}

//3,4,5,6,7
//show as deck.go

//8 go test

//9
func getMap() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	flowers := make(map[string]string)
	flowers["white"] = "#FFFFFF"
	flowers["green"] = "#00FF00"
	flowers["red"] = "#FF0000"
	delete(colors, "red")

	printMap(colors)
	fmt.Println("----------")
	printMap(flowers)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}

//10
func (eb englishBot) getGreeting() string {
	return "Hi there."
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
