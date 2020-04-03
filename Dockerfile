FROM golang:1.13.4-alpine

WORKDIR /app

ENV PORT=5000

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY /api /app/api
COPY main.go .

RUN go build

CMD ["./scg"]