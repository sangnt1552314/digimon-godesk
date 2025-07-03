FROM golang:1.23-alpine 

WORKDIR /app

# Install necessary packages
RUN apk add --no-cache git

# Install Air for hot reloading (compatible with Go 1.23)
RUN go install github.com/cosmtrek/air@v1.51.0

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE ${PORT:-8080}

CMD ["air", "-c", ".air.toml"]