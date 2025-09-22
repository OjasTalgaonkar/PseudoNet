#!/bin/bash

# Kill any previous runs of server or client
pkill -f udp_server.go
pkill -f udp_client.go

# Start the TCP server in the background
echo "Starting UDP server..."
go run Udp_server.go &
SERVER_PID=$!

# Wait a moment for the server to start
sleep 1

# Check if server is running
if ! kill -0 $SERVER_PID 2>/dev/null; then
    echo "Server failed to start. Exiting."
    exit 1
fi

# Start clients
echo "Starting UDP clients..."
go run udp_client.go &
go run udp_client.go &

# Trap Ctrl+C to stop server and clients
trap "echo 'Stopping server and clients...'; kill $SERVER_PID; pkill -f udp_client.go; exit" SIGINT

# Keep the script alive
while true; do
    sleep 1
done
