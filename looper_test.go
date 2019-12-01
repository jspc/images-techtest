package main

import (
	"testing"
)

func TestNewLooper(t *testing.T) {
	count := 3

	l := NewLooper(count)

	if len(l.Pool) != count {
		t.Errorf("expected %d clients, received %d", count, len(l.Pool))
	}
}

func TestLooper_NextClient(t *testing.T) {
	count := 3

	l := NewLooper(count)

	c := l.NextClient()
	if c.ID != 0 {
		t.Errorf("expected 0, received %d", c.ID)
	}

	c = l.NextClient()
	if c.ID != 1 {
		t.Errorf("expected 1, received %d", c.ID)
	}

	l.NextClient()
	l.NextClient()
	l.NextClient()

	c = l.NextClient()
	if c.ID != 2 {
		t.Errorf("expected 2, received %d", c.ID)
	}

}

func TestLooper_Loop(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("unexpected error: %+v", err)
		}
	}()

	count := 1

	l := NewLooper(count)
	l.Pool[0].c = stubDownloader{"some-image", 200, false}

	urls := URLs{"https://example.com/image.png"}

	fC := make(chan []byte, 1)
	fE := make(chan error, 1)

	l.Loop(urls, fC, fE)
}
