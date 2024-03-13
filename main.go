package main

import (
	"fmt"
	"log"
	"os"
)

var debug bool
var towrite string

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments (You must supply one file to edit)")
	}
	fmt.Printf("Welcome to ged! You are editing the file %v \n", os.Args[1])
	debug = false
	ufile := os.Args[1]
	_, err := os.Create(ufile)
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	for towrite != "exit" {
		openedfile, err := os.OpenFile(ufile, os.O_APPEND|os.O_RDWR, os.ModeAppend)
		fmt.Scan(&towrite)
		if towrite == "exit" {
			return
		}
		_, err = openedfile.WriteString(towrite)
		openedfile.WriteString("\n")
		if err != nil {
			log.Fatal("Error writing to file!")
		}
	}
	return
}
