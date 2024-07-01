#!/bin/bash

SERVER_URL="http://localhost:8080"

BLOCK_DATA='{
    "id": "a65e9803bb37256c4a663a5c1b",
    "view": 1234
}'

VOTE_DATA='{
    "block_id": "a65e9803bb37256c4a663a5c1b"
}'

# send block data 

echo "Sending block data..."
BLOCK_RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$SERVER_URL/block" -H "Content-Type: application/json" -d "$BLOCK_DATA")

if [ "$BLOCK_RESPONSE" -eq 200 ]; then
    echo "Block data sent successfully."
else
    echo "Failed to send block data. HTTP status code: $BLOCK_RESPONSE"
fi

# send vote data

echo "Sending vote data..."
VOTE_RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$SERVER_URL/vote" -H "Content-Type: application/json" -d "$VOTE_DATA")

if [ "$VOTE_RESPONSE" -eq 200 ]; then
    echo "Vote data sent successfully."
else
    echo "Failed to send vote data. HTTP status code: $VOTE_RESPONSE"
fi

sleep 1

# show server output
echo "Checking server output..."
LOG_FILE="server_output.log"

# save server log
pgrep -f "go run main.go" | xargs -I {} sh -c "cat /proc/{}/fd/1" > $LOG_FILE

# the log
if grep -q "Accepted Block: ID=a65e9803bb37256c4a663a5c1b, View=1234" $LOG_FILE; then
    echo "Test passed: Block was accepted."
else
    echo "Test failed: Block was not accepted."
fi
