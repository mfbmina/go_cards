package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

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

func deal(d deck, amount int) (deck, deck) {
  hand := d[:amount]
  cards := d[amount:]

  return hand, cards
}

func (d deck) toString() string {
  return strings.Join(d, "\n")
}

func (d deck) saveToFile(filename string) error{
  cardsString := d.toString()
  message := []byte(cardsString)

  return ioutil.WriteFile(filename, message, 0666)
}
