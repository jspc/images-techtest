package main

import (
	"reflect"
	"testing"
)

func TestReadURLs(t *testing.T) {
	for _, test := range []struct {
		name        string
		filename    string
		expect      URLs
		expectError bool
	}{
		{"happy path", "testdata/images.json", URLs{"https://i.pinimg.com/236x/0d/a8/87/0da8872e1ca3e247aef7f75f64a75a5f--learn-coding-logos.jpg"}, false},
		{"missing file", "testdata/nonsuch.json", URLs{}, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := ReadURLs(test.filename)
			if err == nil && test.expectError {
				t.Errorf("expected error")
			}

			if err != nil && !test.expectError {
				t.Errorf("unexpected error: %+v", err)
			}

			if !reflect.DeepEqual(test.expect, got) {
				t.Errorf("expected %+v, received %+v", test.expect, got)
			}

		})
	}
}
