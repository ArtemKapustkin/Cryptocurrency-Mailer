FROM golang:latest

WORKDIR /genesis_test_case

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main main.go

CMD ["./main"]