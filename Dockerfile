FROM golang:1.23-alpine 

WORKDIR /app

# Install necessary packages
RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE ${PORT:-8080}

CMD ["go", "run", "cmd/main.go"]