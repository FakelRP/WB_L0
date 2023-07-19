run-consumer:
	go run ./cmd/order-consumer

run-stream-creator:
	go run ./cmd/stream-creator

install-nats:
	go get github.com/nats-io/nats-server/v2

run-nats-js:
	nats-server -js