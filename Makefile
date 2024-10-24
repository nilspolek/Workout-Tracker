TARGET = tracker

run: build
	@./bin/$(TARGET)

build:
	@mkdir -p bin
	@go build -o bin/$(TARGET) main.go

test:
	go test -v -count=1 ./...

clean:
	@rm -rf bin
