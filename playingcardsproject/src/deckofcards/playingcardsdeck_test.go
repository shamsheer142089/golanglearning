package main

import "testing"

func TestCreateCards(t *testing.T) {
	deck:=createCards()
	if(len(deck)!=40){
		t.Errorf("Length Is Not Matching %v : ", len(deck))
	}
	if deck[0]!="Club Of Aces" {
		t.Errorf("First Card Is Not Matching")
	}
}
