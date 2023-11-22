package dictionary

import (
	"encoding/json"

	"os"
)

// Dictionary struct represents the dictionary.
type Dictionary struct {
	filename string
	entries  map[string]string
	addCh    chan entryOperation
	removeCh chan string
}

type entryOperation struct {
	word       string
	definition string
}

// NewDictionary creates a new Dictionary instance.
func NewDictionary(filename string) *Dictionary {
	d := &Dictionary{
		filename: filename,
		entries:  make(map[string]string),
		addCh:    make(chan entryOperation),
		removeCh: make(chan string),
	}
	d.load()

	go d.listenForOperations()

	return d
}

// Add adds a word and its definition to the dictionary.
func (d *Dictionary) Add(word, definition string) {
	d.addCh <- entryOperation{word, definition}
}

// Get retrieves the definition of a word from the dictionary.
func (d *Dictionary) Get(word string) string {
	return d.entries[word]
}

// Remove deletes a word and its definition from the dictionary.
func (d *Dictionary) Remove(word string) {
	d.removeCh <- word
}

// List returns all the words and their definitions in the dictionary.
func (d *Dictionary) List() map[string]string {
	return d.entries
}

func (d *Dictionary) load() {
	data, err := os.ReadFile(d.filename)
	if err != nil {

		return
	}

	err = json.Unmarshal(data, &d.entries)
	if err != nil {
		panic(err)
	}
}

func (d *Dictionary) save() {
	data, err := json.Marshal(d.entries)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(d.filename, data, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func (d *Dictionary) listenForOperations() {
	for {
		select {
		case add := <-d.addCh:
			d.entries[add.word] = add.definition
			d.save()
		case remove := <-d.removeCh:
			delete(d.entries, remove)
			d.save()
		}
	}
}
