clean:
	rm -f bin/gocastle

build: clean
	go build -o bin/gocastle

run:
	bin/gocastle

buildrun: build run