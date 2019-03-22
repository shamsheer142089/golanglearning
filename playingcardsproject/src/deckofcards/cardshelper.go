package main

import "strings"

func (deck deckOfCards) toString() string {
	cards:=[]string(deck)
	data:=strings.Join(cards,",")
	return data;
}

func toByteSlice(data string) []byte{
	dataInString:=[]byte(data)
	return dataInString
}

func (deck deckOfCards) toByteSlice() []byte{
	dataInStringFormat:=deck.toString()
	return toByteSlice(dataInStringFormat)
}

