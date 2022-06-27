package main

import (
	"log"
	"os"
)

func main() {
	const directory = "assets"
	log.Print("1. Create a directory named \"assets\" if it does not exist")

	err := os.MkdirAll(directory, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}
	log.Print("1. Directory named \"assets\" created")
}
