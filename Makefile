BUILD_DIRECTORY = _build

client:
	go build -o cmd/client/${BUILD_DIRECTORY}/whisper ./cmd/client

server:
	go build -o cmd/server/${BUILD_DIRECTORY}/whisper-server ./cmd/server

clean-client:
	rm -rf cmd/client/${BUILD_DIRECTORY}

clean-server:
	rm -rf cmd/server/${BUILD_DIRECTORY}

run-client:
	go run ./cmd/client/main.go

run-server:
	go run ./cmd/server/main.go

all: clean-client client clean-server server

.PHONY: client server clean run-client run-server