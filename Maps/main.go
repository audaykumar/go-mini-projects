package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// colors := make(map[string]string)

	colors["white"] = "#ffffff"

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("Hex code for %s is %s\n", color, hex)
	}
}
