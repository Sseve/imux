##

.PHONY: clean, build

build:
	go build -trimpath -ldflags="-s -w -X main.version=0.1.1" -o example

clean:
	rm -fr example