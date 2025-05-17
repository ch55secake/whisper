BUILD_DIRECTORY = _build

client:
	go build -o cmd/client/${BUILD_DIRECTORY}/whisper ./cmd/client

server:
	go build -o cmd/server/${BUILD_DIRECTORY}/whisper-server ./cmd/server

clean:
	rm -rf cmd/client/${BUILD_DIRECTORY} && rm -rf cmd/server/${BUILD_DIRECTORY}

run_client:
	go run ./cmd/client/main.go

run_server:
	./cmd/server/_build/whisper-server

all: clean client server

.PHONY: client server clean run_client run_server