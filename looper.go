package main

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

type Looper struct {
	Pool []Client

	sem  *semaphore.Weighted
	c    int
	lock sync.Mutex
}

func NewLooper(count int) (l Looper) {
	l.Pool = make([]Client, count)

	l.sem = semaphore.NewWeighted(int64(count))
	l.c = 0
	l.lock = sync.Mutex{}

	for i := 0; i < count; i++ {
		l.Pool[i] = NewClient(i)
	}

	return
}

func (l *Looper) NextClient() Client {
	l.lock.Lock()
	defer l.lock.Unlock()

	c := l.Pool[l.c%len(l.Pool)]

	l.c += 1

	return c
}

func (l Looper) Loop(urls URLs, files chan []byte, errors chan error) {
	ctx := context.Background()

	for _, u := range urls {
		l.sem.Acquire(ctx, 1)

		go func(url string) {
			defer l.sem.Release(1)

			client := l.NextClient()
			file, err := client.Download(url)

			files <- file
			errors <- err

		}(u)
	}

	// Block until complete
	l.sem.Acquire(ctx, int64(len(l.Pool)))
}
