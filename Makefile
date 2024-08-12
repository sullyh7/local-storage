run: build
	@./bin/localstorage

initdb:
	@go run scripts/main.go

css:
	@npx tailwindcss -i view/css/app.css -o public/styles.css

templ:
	@templ generate --watch --proxy=http://localhost:3000

build: css
	@templ generate view
	@go build -tags prod -o bin/localstorage main.go 

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@templ generate
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss

