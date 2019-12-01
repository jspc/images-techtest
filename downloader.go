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
	c DownloadClient
}

func NewClient() Client {
	return Client{
		c: &http.Client{},
	}
}

func (c Client) Download(url string) (data []byte, err error) {
	resp, err := c.c.Get(url)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = DownloadError{status: resp.StatusCode}

		return
	}

	return ioutil.ReadAll(resp.Body)
}
