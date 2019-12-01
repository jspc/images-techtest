package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestNewStorage(t *testing.T) {
	dir := "testdata/storage_test"

	err := os.RemoveAll(dir)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)

		t.Fail()
	}

	_, err = NewStorage(dir)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)

		t.Fail()
	}

	_, err = os.Stat(dir)
	if err != nil {
		t.Errorf("Expected dir %q does not exist", dir)
	}
}

func TestStorage_Save(t *testing.T) {
	dir := "testdata/storage_test"
	fn := "some_file.txt"

	full := filepath.Join(dir, fn)

	s, err := NewStorage(dir)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)

		t.Fail()
	}

	contents := "hello, world!"
	f := File{filename: fn, contents: []byte(contents)}

	err = s.Save(f)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)

		t.Fail()
	}

	received, err := ioutil.ReadFile(full)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)

		t.Fail()
	}

	if contents != string(received) {
		t.Errorf("expected %q, received %q", contents, string(received))
	}
}
