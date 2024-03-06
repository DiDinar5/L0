FROM nats-streaming:latest

EXPOSE 4222 8222

CMD ["-m" "8222"]