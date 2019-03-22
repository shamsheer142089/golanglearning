package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a slice of cards
type deckOfCards []string


func (deck deckOfCards) printCards(){
	for index,card:= range deck{
		fmt.Println(index,"~~~~~~",card)
	}
}

func createCards() deckOfCards{
	cards:=deckOfCards{}
	suits := []string {"Club","Spades","Diamond","Hearts"}
	cardValues := []string {"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten"}

	for _,suit := range suits{
		for _,card := range cardValues {
			cards=append(cards,card + " Of " + suit)
		}
	}
	return cards
}

func deal(dec deckOfCards,size int) (deckOfCards, deckOfCards){
	return dec[:size],dec[size:]
}

func (cards deckOfCards) writeDataToFile(fileName string){
	ioutil.WriteFile(fileName,cards.toByteSlice(),0644)
}

func getDataFromFile(fileName string) deckOfCards{
	byteslice,error:=ioutil.ReadFile(fileName)
	if(error!=nil){
		fmt.Println("Exception occured while reading the file",error)
		os.Exit(1)
		}
	dataInStringFormat:=string(byteslice)
	sliceOfString:=strings.Split(dataInStringFormat,",")
	deck:=deckOfCards(sliceOfString)
	return deck

}

func (deck deckOfCards) shuffleCards(){
	seedValue:=rand.NewSource(time.Now().UnixNano())
	rnd:=rand.New(seedValue)
	for idx:=range deck{
		random:=rnd.Intn(len(deck)-1)
		deck[idx],deck[random]=deck[random],deck[idx]
	}
}