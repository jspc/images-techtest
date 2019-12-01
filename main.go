package main

import (
	"fmt"
	"log"
	"os"
)

var (
	args      = os.Args
	outputDir = "./data"
	workers   = 3

	incorrectUsageError = fmt.Errorf(`Usage:
    %s images.json

Where:
    images.json: path to json file containing a list of images to download
`, args[0])
)

func main() {
	if len(args) != 2 {
		panic(incorrectUsageError)
	}

	err := realmain(args[1])
	if err != nil {
		log.Fatal(err)
	}
}
