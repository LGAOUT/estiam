package dictionary

type Dictionary map[string]string

// Add
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// Get
func (d Dictionary) Get(word string) string {
	return d[word]
}

// Remove
func (d Dictionary) Remove(word string) {
	delete(d, word)
}

// List
func (d Dictionary) List() map[string]string {
	return d
}
