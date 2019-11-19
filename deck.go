package main

import (
  "fmt"
  "io/ioutil"
  "math/rand"
  "strings"
  "time"
  "os"
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

func deal(d deck, amount int) (deck, deck) {
  hand := d[:amount]
  cards := d[amount:]

  return hand, cards
}

func (d deck) print() {
  for _, card := range d {
    fmt.Println(card)
  }
}

func (d deck) toString() string {
  return strings.Join(d, "\n")
}

func (d deck) saveToFile(filename string) error {
  cardsString := d.toString()
  message := []byte(cardsString)

  return ioutil.WriteFile(filename, message, 0666)
}

func deckFromFile(filename string) deck {
  bs, err := ioutil.ReadFile(filename)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  cards := strings.Split(string(bs), "\n")

  return deck(cards)
}

func (d deck) shuffle() {
  source := rand.NewSource(time.Now().UnixNano())
  r := rand.New(source)

  for i := range d {
    ni := r.Intn(len(d) - 1)
    d[i], d[ni] = d[ni], d[i]
  }
}
