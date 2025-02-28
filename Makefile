all: build

.PHONY: build

build:
	go build -o w.exe .

run: build
	.\w.exe go