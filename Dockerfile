FROM golang:1.23-alpine 

WORKDIR /app

# Install necessary packages
RUN apk add --no-cache git

RUN go get github.com/joho/godotenv

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE ${PORT:-8080}

CMD ["go", "run", "main.go"]