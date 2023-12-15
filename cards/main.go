package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades" //use := for the every first initialization, let golang identify the variable type
	card := newCard()
	
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
