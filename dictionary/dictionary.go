package dictionary

type Dictionary map[string]string

// Add adds a word and its definition to the dictionary.
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// Get retrieves the definition of a word from the dictionary.
func (d Dictionary) Get(word string) string {
	return d[word]
}

// Remove deletes a word and its definition from the dictionary.
func (d Dictionary) Remove(word string) {
	delete(d, word)
}

// List returns all the words and their definitions in the dictionary.
func (d Dictionary) List() map[string]string {
	return d
}
