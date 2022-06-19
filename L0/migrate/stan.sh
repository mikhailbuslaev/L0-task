echo "Running nats streaming container...";
docker run -p 4222:4222 -p 8222:8222 -ti nats-streaming:latest;