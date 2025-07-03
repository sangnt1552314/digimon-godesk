FROM golang:1.23-alpine 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

EXPOSE 8080

CMD ["/app/main"]