.PHONY: all
all:
	make build
	make test
	make pack

.PHONY: build
build:
	cargo build --target wasm32-unknown-unknown --release

.PHONY: test
test:
	cargo test

.PHONY: protogen
protogen:
	substreams protogen --exclude-paths sf/substreams,google

.PHONY: pack
pack:
	substreams pack