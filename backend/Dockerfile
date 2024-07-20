FROM golang:1.22.4

WORKDIR /app

COPY . .

RUN apt-get update

RUN go mod download
RUN go build -o blog-app ./app/cmd/main.go

EXPOSE 8080

CMD ["./blog-app"]
