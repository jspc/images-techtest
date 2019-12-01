# images tech test

As per: https://gist.github.com/ypapax/de5a83870273c2e7cc8dc1ea62c81fe1

## Testing

A convenience `Make` target exists in this project to test with coverage:

```bash
$ make test
```

Otherwise:

```bash
$ go test
```

## Building

A convenience `Make` target exists to build the binary as `./demo`, as per the spec:

```bash
$ make demo
```

## Usage:

```bash
$ ./demo
panic: Usage:
    ./demo images.json

Where:
    images.json: path to json file containing a list of images to download


```

The binary will panic if it does not find the arguments it expects.
