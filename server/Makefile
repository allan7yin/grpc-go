.PHONY: generate
generate:
	cd proto; buf generate

.PHONY: run
run:
	cd cmd; go run main.go

.PHONY: clean
clean:
	rm -rf pkg

.PHONY: build
build:
	cd cmd; go build main.go