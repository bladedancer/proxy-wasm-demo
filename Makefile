.PHONY: build.docker lint

all: build.docker run

build.docker:
	docker run -it -w /tmp/proxy_wasm_demo -v $(shell pwd):/tmp/proxy_wasm_demo tinygo/tinygo:0.16.0 \
		tinygo build -o /tmp/proxy_wasm_demo/main.go.wasm -scheduler=none -target=wasi \
		/tmp/proxy_wasm_demo/main.go

run:
	getenvoy run standard:1.17.0 -- --config-path ./envoy.yaml --concurrency 2 --log-format '%v'