
all: build

build:
	@echo -e "Building..."
	@go build -o bin/templ .

dev:
	@echo -e "Starting development server..."
	tailwindcss -i ./global.css -o ./view/public/styles.css --watch=always &
	~/go/bin/templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."


.PHONY: all dev build