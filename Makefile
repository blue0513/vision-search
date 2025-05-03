APP_NAME=vision-search

.PHONY: build install clean

build:
	go build -o $(APP_NAME) main.go

install:
	go install

clean:
	rm -f $(APP_NAME)
