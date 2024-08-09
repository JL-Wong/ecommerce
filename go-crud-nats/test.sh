#!/bin/bash

# URL of your REST API that publishes to NATS
REST_API_URL="http://localhost:8000/users"

# Number of requests to send
NUM_REQUESTS=100

# Number of concurrent requests
CONCURRENT_REQUESTS=10

# Function to send a request
send_request() {
  i=$1
  curl -s -X GET $REST_API_URL -d "{\"subject\":\"user.get\", \"message\":\"user001$i\"}" -H "Content-Type: application/json" >/dev/null
}

export -f send_request
export REST_API_URL

# Capture the start time
start_time=$(date +%s)

# Use GNU parallel to send requests concurrently
seq 1 $NUM_REQUESTS | parallel -j $CONCURRENT_REQUESTS send_request

# Capture the end time
end_time=$(date +%s)

# Calculate the elapsed time
elapsed_time=$((end_time - start_time))

# Display the elapsed time
echo "Time taken to complete the requests: $elapsed_time seconds"

read -p "Press [Enter] key to exit..."
