package main

import "fmt"

func main() {
	allCards:=createCards()
	allCards.printCards()
	inHandCards,remainingCards :=deal(allCards,3)
	inHandCards.printCards()
	remainingCards.printCards()
	data:=inHandCards.toString()
	fmt.Println(data)
	fmt.Println(toByteSlice(data))
	allCards.writeDataToFile("CardList.txt")
	//get the cards from file
	cardsFromfile:=getDataFromFile("CardList.txt")
	fmt.Println("Data From File")
	cardsFromfile.printCards()


	//shuffle cards
	cardsFromfile.shuffleCards();

	//getDataFromFile("NotAvailable.txt")



}
