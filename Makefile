.PHONY: all build test clean install ontology auth auth-clear run

all: build

build:
	go build -o nlm ./cmd/nlm

install:
	go install ./cmd/nlm

test:
	go test ./...

clean:
	rm -f nlm

generate:
	cd proto && go tool buf generate
ontology:
	go run ./cmd/podcast-ontology podcast/config.yaml

auth: build
	./nlm auth
	scripts/secrets.sh push

auth-clear:
	scripts/secrets.sh clear

run: build
	scripts/secrets.sh exec ./nlm $(ARGS)
