build:
	go build -o bin/gocastle

run:
	bin/gocastle

buildrun: build run