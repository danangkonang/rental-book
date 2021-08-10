FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

# RUN go build -o binary
RUN go build main.go

EXPOSE 9000

ENTRYPOINT ["/app/main"]