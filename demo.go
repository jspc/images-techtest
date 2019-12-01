package main

func realmain(inputFile string) (err error) {
	urls, err := ReadURLs(inputFile)
	if err != nil {
		return
	}

	storage, err := NewStorage(outputDir)
	if err != nil {
		return
	}

	looper := NewLooper(workers, urls, storage)

	errsChan := make(chan error, 1)
	go func() {
		for err := range errsChan {
			if err != nil {
				return
			}
		}
	}()

	looper.Loop(errsChan)

	return
}
