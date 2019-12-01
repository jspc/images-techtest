package main

import (
	"log"
)

func realmain(inputFile string) (err error) {
	urls, err := ReadURLs(inputFile)
	if err != nil {
		return
	}

	storage, err := NewStorage(outputDir)
	if err != nil {
		return
	}

	looper := NewLooper(workers, urls, storage)

	errsChan := make(chan error, 1)
	go func() {
		for err := range errsChan {
			if err != nil {
				log.Printf("%+v", err)
			}
		}
	}()

	looper.Loop(errsChan)

	return
}
