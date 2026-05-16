FROM golang:1.26.2

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o server ./cmd/server

EXPOSE 8080

CMD ["./server"]