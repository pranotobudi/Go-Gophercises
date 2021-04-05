package main

import deck "github.com/pranotobudi/Go-Gophercises/Go-Deck/deck"

func main() {
	cards := deck.New(deck.DefaultSort)
	deck.Shuffle(cards)
}
