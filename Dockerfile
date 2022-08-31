FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod init web
RUN go mod tidy
RUN go build -o main .

CMD ["./main"]

EXPOSE 8000
