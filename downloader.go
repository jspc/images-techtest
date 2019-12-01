package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type DownloadError struct {
	expect *int
	status int
}

func (d DownloadError) Error() string {
	expect := 200
	if d.expect != nil {
		expect = *d.expect
	}

	return fmt.Sprintf("Unexpected status code %d. Expected: %d", d.status, expect)
}

type DownloadClient interface {
	Get(string) (*http.Response, error)
}

type Client struct {
	ID int
	c  DownloadClient
}

func NewClient(id int) Client {
	return Client{
		ID: id,
		c:  &http.Client{},
	}
}

func (c Client) Download(url string) (data []byte, err error) {
	c.Print("Downloading", url)

	resp, err := c.c.Get(url)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = DownloadError{status: resp.StatusCode}

		return
	}

	c.Print("Completed", url)

	return ioutil.ReadAll(resp.Body)
}

func (c Client) Print(operation, url string) {
	fmt.Printf("#%d - %s %s\n", c.ID, operation, url)
}
