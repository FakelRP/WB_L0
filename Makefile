run-consumer:
	go run ./cmd/order-consumer

run-stream-creator:
	go run ./cmd/stream-creator

install-nats:
	go get github.com/nats-io/nats-server/v2

run-nats-js:
	nats-server -js

run-docker-js:
	sudo docker run -p 4222:4222 -ti nats:latest -js

start-postgres-docker:
	sudo docker start 4c44ccdc3187