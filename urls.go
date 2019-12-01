package main

import (
	"encoding/json"
	"io/ioutil"
)

type URLs []string

func ReadURLs(filename string) (u URLs, err error) {
	u = make(URLs, 0)

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = json.Unmarshal(f, &u)

	return
}
