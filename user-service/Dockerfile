FROM debian:latest

RUN mkdir /app
WORKDIR /app
ADD user-service /app

ENV PORT=:8000

CMD ["./packet-service"]