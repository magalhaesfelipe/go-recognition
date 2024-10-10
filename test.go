package main

import (
	"fmt"
	"sort"
)

type Book struct {
	Title string
	Pages int
}

func sliceOfStructs() {
	
	bookPages := []Book{
		{"Lord of the Rings", 1000},
		{"Duna", 700},
		{"Harry Potter", 600},
		{"Lord of the Flies", 224},
	}

	// Print books in the order they were inserted
	fmt.Println("Books in insertion order:")
	for _, book := range bookPages {
		fmt.Printf("%s: %d pages\n", book.Title, book.Pages)
	}

	// Sort the slice alphabetically by title
	sort.Slice(bookPages, func(i, j int) bool {
		return bookPages[i].Title < bookPages[j].Title
	})

	fmt.Println("\nBooks sorted by title:")
	for _, book := range bookPages {
		fmt.Printf("%s: %d pages\n", book.Title, book.Pages)
	}


} 