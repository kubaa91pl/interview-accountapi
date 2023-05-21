FROM golang:alpine

WORKDIR /accountclient

COPY go.mod ./
RUN go mod download
COPY . .

CMD go test ./... -v -cover