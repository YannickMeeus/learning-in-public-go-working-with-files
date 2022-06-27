package main

import (
	"log"
	"os"
)

func main() {
	const directory = "assets"
	log.Print("1. Create a directory named \"assets\" if it does not exist")

	err := os.Mkdir(directory, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}
}
