package dictionary

import (
	"encoding/json"

	"os"
)

// Dictionary struct represents the dictionary.
type Dictionary struct {
	filename string
	entries  map[string]string
}

// NewDictionary creates a new Dictionary instance.
func NewDictionary(filename string) *Dictionary {
	d := &Dictionary{filename: filename}
	d.load()
	return d
}

// Add adds a word and its definition to the dictionary.
func (d *Dictionary) Add(word, definition string) {
	d.entries[word] = definition
	d.save()
}

// Get retrieves the definition of a word from the dictionary.
func (d *Dictionary) Get(word string) string {
	return d.entries[word]
}

// Remove deletes a word and its definition from the dictionary.
func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
	d.save()
}

// List returns all the words and their definitions in the dictionary.
func (d *Dictionary) List() map[string]string {
	return d.entries
}

func (d *Dictionary) load() {
	data, err := os.ReadFile(d.filename)
	if err != nil {
		d.entries = make(map[string]string)
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
