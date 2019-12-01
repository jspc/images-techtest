package main

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

type Looper struct {
	Pool    []Client
	URLs    URLs
	Storage Storage

	sem  *semaphore.Weighted
	c    int
	lock sync.Mutex
}

func NewLooper(count int, urls URLs, storage Storage) (l Looper) {
	l.Pool = make([]Client, count)
	l.URLs = urls
	l.Storage = storage

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

func (l Looper) Loop(errors chan error) {
	ctx := context.Background()

	for _, u := range l.URLs {
		l.sem.Acquire(ctx, 1)

		go func(url string) {
			defer l.sem.Release(1)

			client := l.NextClient()
			file, err := client.Download(url)
			if err != nil {
				errors <- err

				return
			}

			errors <- l.Storage.Save(file)

		}(u)
	}

	// Block until complete
	l.sem.Acquire(ctx, int64(len(l.Pool)))
}
