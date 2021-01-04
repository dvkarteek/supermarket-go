docker: build
	docker build . --rm -t supermarket-service:latest
build:
	GOOS=linux GOARCH=amd64 go build .
	 