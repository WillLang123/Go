package main

//go run main.go deck.go

func main() {
	// var card string = "ace of spades"
	// var card = "ace of spades"
	/* card := "ace of pades"
	card = "ace of spades"
	card = newCard() */
	cards := deck{newCard(), newCard()}
	cards = append(cards, newCard())

	cards.print()

}

func newCard() string {
	return "five of diamonds"
}
