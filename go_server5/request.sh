#!/bin/bash

URL="http://localhost:8081/hello"
THREADS=5
REQ_PER_THREAD=2000    # 10000 / 5

make_requests() {
  for i in $(seq 1 $REQ_PER_THREAD); do
    curl -s -o /dev/null "$URL"
    sleep 0.001    # 1 millisecond delay
  done
  echo "Thread finished $REQ_PER_THREAD requests"
}

echo "Starting load test (10000 requests, 5 threads, 1ms delay each)..."

for t in $(seq 1 $THREADS); do
  make_requests &
done

wait
echo "Completed 10000 requests with 5 threads."
