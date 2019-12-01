package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	defer func() {
		err := recover()

		if err != nil {
			t.Errorf("unexpected error: %+v", err)
		}
	}()

	args = []string{"go.test", "testdata/images.json"}
	outputDir = "testdata/main_test"

	main()
}

func TestMain_BadArgs(t *testing.T) {
	defer func() {
		err := recover()

		if err == nil {
			t.Errorf("expected error, received none")
		}
	}()

	args = []string{"go.test"}
	outputDir = "testdata/main_test"

	main()
}
