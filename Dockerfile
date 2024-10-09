FROM golang:1.22.4

WORKDIR /usr/wallets

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make build

# CMD ["./bin/server", "serve"]

