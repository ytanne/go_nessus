FROM golang:1.17.0-buster
WORKDIR /app
COPY . .
RUN go build -ldflags "-w -s" -o tg_bot cmd/main.go
RUN apt update && apt install -y nmap
CMD ["/app/tg_bot"]