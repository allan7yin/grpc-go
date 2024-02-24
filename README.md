<img src="https://github.com/allan7yin/grpc-go/assets/66652405/4ea42bce-c7a4-4d4c-b381-6b056e93078b" alt="What is gRPC" width="400" />

# gRPC 
This is a quick spinup of a gRPC client and server in Go. It includes examples of different types of RPCs including Unary, Server Streaming, Client Streaming, and Bidirectional Streaming.

## Features
- Unary RPCs where the client sends a single request to the server and gets a single response back.
- Server streaming RPCs where the client sends a request to the server and gets a stream to read a sequence of messages back.
- Client streaming RPCs where the client writes a sequence of messages and sends them to the server, again using a provided stream.
- Bidirectional streaming RPCs where both sides send a sequence of messages using a read-write stream.

## Running the Code

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. To run the server, navigate to the `server` directory and run `make run`.
4. To run the client, navigate to the `client` directory and run `make run` in a separate terminal session.

If changes are made to the message and service definitions, re-compile protocol buffers via `make generate`
## Dependencies

This project uses the following dependencies:
- [google.golang.org/grpc](https://pkg.go.dev/google.golang.org/grpc)
- [google.golang.org/protobuf](https://pkg.go.dev/google.golang.org/protobuf)

## Next
Look to dockerize this application -> implement this logic in microservices when deploying plant-trait model
