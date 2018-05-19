package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#4BF745",
		"white": "#FFFFFF",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%s: %s\n", color, hex)
	}
}
