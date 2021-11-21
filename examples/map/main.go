package main

import "fmt"

func main() {
	//method one
	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"green": "#00FF00",
	// }

	//method two
	//var colors map[string]string

	//method three
	colors := make(map[string]string)
	colors["white"] = "#FFFFFF"
	colors["green"] = "#00FF00"
	colors["red"] = "#FF0000"

	fmt.Println(colors)

	//delete(colors, "green")
	//fmt.Print(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
