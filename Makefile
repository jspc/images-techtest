test:
	# TODO: use mktemp for file name
	go test -covermode=count -coverprofile=count.out -v
	go tool cover -html=count.out


demo:
	go build -o demo
