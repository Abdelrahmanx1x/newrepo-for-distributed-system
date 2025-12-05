This project separates server and client into different directories so each can be built as its own executable.

Build server:
    cd server
    go build -o server

Build client:
    cd client
    go build -o client

Run:
    ./server/server
    ./client/client localhost:8080
