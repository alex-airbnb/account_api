
FROM golang:1.15

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]