// demonstrates objects in Go
// Name: Neel Ranawat
// KSUID: 000988101
// Class: CS 4308 W03 Spring 2025
package main

import "fmt"

type Animal struct {
	Name           string
	ScientificName string
	FavoriteArtist string
}

func (p Animal) Description() {
	fmt.Printf("The %s's scientific name is %s and it loves listening to %s\n", p.Name, p.ScientificName, p.FavoriteArtist)
}

func main() {
	person := Animal{Name: "Wolf", ScientificName: "Canis Lupus", FavoriteArtist: "Nina Simone"}
	person.Description()
}
