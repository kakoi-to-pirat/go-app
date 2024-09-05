.PHONY: build
build:
	templ generate internal/view
	go build -o ./bin/app cmd/app/main.go

.PHONY: run
run: build
	./bin/app