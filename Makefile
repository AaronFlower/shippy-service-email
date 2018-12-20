build:
	GOOS=linux GOARCH=amd64 go build -o service.email
	docker build --rm -t service.email .

run:
	docker run --net="host" \
		-p 50054 \
		-e MICRO_SERVER_ADDRESS=:50054 \
		-e MICRO_REGISTRY=mdns \
		service.email

clean:
	go clean
	rm service.email
