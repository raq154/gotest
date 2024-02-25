package main

import (
	"os"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	testDeck := newDeck()
	if len(testDeck) != 40 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", len(testDeck), 40)
	}
	firstCard := testDeck[0]
	if firstCard != "Ace of Spades" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", firstCard, "Ace of Spades")
	}

	lastCard := testDeck[len(testDeck)-1]
	if lastCard != "Nine of Clubs" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", firstCard, "Nine of Clubs")
	}
}

func TestSaveToFile(t *testing.T) {
	tmpdir := t.TempDir()
	file, err := os.CreateTemp(tmpdir, "*")
	if err != nil {
		t.Error("Unable to create file, cannot continue test")
	}

	testDeck := newDeck()
	if err := testDeck.saveToFile(file.Name()); err != nil {
		t.Errorf("unexpected error in saving file %s", err)
	}
	defer os.Remove(file.Name())
}

func TestReadFromFile(t *testing.T) {
	tmpdir := t.TempDir()
	file, err := os.CreateTemp(tmpdir, "*")
	if err != nil {
		t.Errorf("Unable to create file, cannot continue test %s", err)
	}

	testDeck := newDeck()
	testDeck.shuffle()
	if err := testDeck.saveToFile(file.Name()); err != nil {
		t.Errorf("unexpected error in saving file %s", err)
	}

	deckFromFile, err := readFromFile(file.Name())
	if err != nil {
		t.Errorf("unexpected error in reading file %s", err)
	}

	if !reflect.DeepEqual(deckFromFile, testDeck) {
		t.Errorf("Decks not equal, got: %s, want: %s.", deckFromFile.toString(), testDeck.toString())
	}

	//cleanup
	defer os.Remove(file.Name())
}
