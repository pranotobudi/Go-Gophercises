package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})
	//Output:
	// Ace of Heart
	// Two of Spade
	// Nine of Diamond
	// Jack of Club
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Suit: Spade, Rank: Ace}
	if cards[0] != exp {
		t.Error("expected Ace of Spade, but received:", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	//shuffleRandSource is a variable as a source of random at card.go
	shuffleRandSource = rand.New(rand.NewSource(0))
	//the output of same input like 0, 1, etc always consistent the same
	//the output of 0, we can use go playground to know the result of rand.NewSource
	//check the slice result for r.rand(52)
	original := New() //no shuffle
	first := original[40]
	second := original[35]

	cards := New(Shuffle)

	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, received %s.", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the second card to be %s, received %s.", second, cards[1])
	}
}
func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("expected 3 Jokers, but received: ", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == 2 || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards, received %d cards", 13*4*3, len(cards))
	}
}
