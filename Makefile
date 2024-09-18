run:
	air & npm run watch --prefix internal/views/web/assets/

setup:
	npm install --prefix internal/views/web/assets/
	go install github.com/air-verse/air@latest

build:
	npm run build --prefix internal/views/web/assets/
	go build .
