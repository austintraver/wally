package main

import (
	"log"

	"github.com/austintraver/wally/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
