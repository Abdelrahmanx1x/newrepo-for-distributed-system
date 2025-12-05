# Realtime Go Chat (Broadcast, no RPC)
# Abdelrahman Mohamed Abdelrahman


Features:
- When a client joins, the server broadcasts: `User [ID] joined` (other clients see it; the joining client does not receive their own join message).
- When a client sends a message, the server broadcasts it to **all other clients** (no self-echo).
- Uses goroutines and buffered channels for concurrent send/receive.
- `clients` map is protected with a `sync.Mutex`.

## Files
- `server.go` — chat server
- `client.go` — simple terminal client
- `Dockerfile.server` — Dockerfile to build the server image
- `Dockerfile.client` — Dockerfile to build the client image
- `go.mod` — module file
- `README.md` — this file

## Build & Run (locally)
Requires Go 1.20+.

Build server:
```bash
go build -o server ./server.go
```

Build client:
```bash
go build -o client ./client.go
```

Run server:
```bash
./server
```

Open multiple terminals and run client:
```bash
./client localhost:8080
```
Type messages in any client; other clients will see them. The client that typed the message will not receive it back (no self-echo).

## Docker
Build server image:
```bash
docker build -f Dockerfile.server -t realtime-go-chat-server .
```

Run:
```bash
docker run -p 8080:8080 realtime-go-chat-server
```

Build client image (optional):
```bash
docker build -f Dockerfile.client -t realtime-go-chat-client .
```

Run client (example, using docker networking to connect to host):
```bash
docker run -it --rm realtime-go-chat-client localhost:8080
```

