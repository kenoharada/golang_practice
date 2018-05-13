package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "ff0000",
		"green": "ff0001",
	}
	delete(colors, "red")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}

}
