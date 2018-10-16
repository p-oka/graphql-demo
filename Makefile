.PHONY: build
build: assets
	go build -o app ./

.PHONY: up
up: build
	./app

.PHONY: assets
assets:
	npm run build
	go-assets-builder -s "/dist" -o assets.go dist/
