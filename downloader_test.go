package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type stubDownloader struct {
	response string
	status   int
	error    bool
}

func (s stubDownloader) Get(string) (resp *http.Response, err error) {
	if s.error {
		err = fmt.Errorf("an error")
	}

	body := ioutil.NopCloser(strings.NewReader(s.response))

	return &http.Response{
		StatusCode: s.status,
		Body:       body,
	}, err
}

func TestClient_Download(t *testing.T) {
	url := "http://example.com/image.jpg"

	for _, test := range []struct {
		name           string
		downloadClient DownloadClient
		expect         string
		expectError    bool
	}{
		{"happy path", stubDownloader{"some-image-data", 200, false}, "some-image-data", false},
		{"erroring http call", stubDownloader{"", 0, true}, "", true},
		{"non-200", stubDownloader{"", 404, false}, "", true},
	} {
		t.Run(test.name, func(t *testing.T) {
			c := Client{c: test.downloadClient}

			got, err := c.Download(url)

			if err == nil && test.expectError {
				t.Errorf("expected error")
			}

			if err != nil && !test.expectError {
				t.Errorf("unexpected error: %+v", err)
			}

			gotString := string(got.contents)
			if gotString != test.expect {
				t.Errorf("expected %q, received %q", gotString, test.expect)
			}
		})
	}
}
