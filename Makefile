.PHONY: build
build:
	go build -o ./bin/app cmd/app/main.go

.PHONY: run
run: build
	./bin/app