.PHONY: all
all: install compile

.PHONY: compile
compile:
	go build
	npm run build

.PHONY: install
install:
	go install

.PHONY: test
test:
	go test ./...

watch:
	dev_appserver.py app.yaml