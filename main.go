package main

import (
	"estiam/dictionary"

	"fmt"
)

func main() {
	// Create a new dictionary
	myDictionary := make(dictionary.Dictionary)

	// Add words and their definitions
	myDictionary.Add("niece", "a daughter of one's brother or sister")
	myDictionary.Add("stair", "a set of steps leading from one floor of a building to another")
	myDictionary.Add("chair", "a separate seat for one person")
	myDictionary.Add("loan", "a sum of money that is expected to be paid back with interest")

	// Get and print a definition
	fmt.Println("The Definition of 'chair':", myDictionary.Get("chair"))

	// List all words and their values
	fmt.Println("All words and definitions:", myDictionary.List())

	// Remove func
	myDictionary.Remove("chair")

	// List without 'chair'
	fmt.Println("All words and definitions after removal:", myDictionary.List())
}
