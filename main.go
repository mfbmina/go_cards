package main

func main() {
  cards := newDeck()

  hand, remainingDeck := deal(cards, 4)

  hand.print()
  remainingDeck.saveToFile("test.txt")
  remainingDeck.print()
}
