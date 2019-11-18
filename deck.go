package main

import "fmt"

type deck []string

func newDeck() deck {
  cards := deck{}
  cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
  cardValues := []string{"Ace", "Two", "Three", "Four"}

  for _, suit := range cardSuits {
    for _, value := range cardValues {
      card := value + " of " + suit
      cards = append(cards, card)
    }
  }

  return cards
}

func (d deck) print() {
  for _, card := range d {
    fmt.Println(card)
  }
}
