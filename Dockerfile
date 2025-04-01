FROM golang:tip-bookworm

WORKDIR /app

RUN go install github.com/air-verse/air@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
