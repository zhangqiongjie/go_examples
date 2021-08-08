package main

import "fmt"

func main() {
	cards := []string{"abcd", printState()}
	cards = append(cards, "hello kevin")
	for a, card := range cards {
		fmt.Println(a, card)
	}
}

func printState() string {
	return "successed"
}
