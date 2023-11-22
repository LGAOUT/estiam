package main

import (
	"estiam/dictionary"

	"fmt"

	"log"

	"net/http"
)

func main() {
	// Create a new dictionary
	myDictionary := dictionary.NewDictionary("dictionary.json")

	// Handle API requests
	myDictionary.HandleRequests()

	// Start the server
	port := 8080
	fmt.Printf("Starting server on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

// 	// Add words and their definitions
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		myDictionary.Add("stair", "a set of steps leading from one floor of a building to another")
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		myDictionary.Add("loan", "a sum of money that is expected to be paid back with interest")
// 	}()

// 	myDictionary.Add("niece", "a daughter of one's brother or sister")
// 	myDictionary.Add("stair", "a set of steps leading from one floor of a building to another")
// 	myDictionary.Add("chair", "a separate seat for one person")
// 	myDictionary.Add("loan", "a sum of money that is expected to be paid back with interest")

// 	// Remove a word
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		myDictionary.Remove("stair")
// 	}()

// 	// Wait for all operations to complete
// 	wg.Wait()

// 	// List all words and their definitions
// 	fmt.Println("All words and definitions:", myDictionary.List())
// }

// 	// Get and print a definition
// 	fmt.Println("The Definition of 'chair':", myDictionary.Get("chair"))

// 	// List all words and their values
// 	fmt.Println("All words and definitions:", myDictionary.List())

// 	// Remove func
// 	myDictionary.Remove("chair")

// 	// List without 'chair'
// 	fmt.Println("All words and definitions after removal:", myDictionary.List())
// }
