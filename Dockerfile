FROM golang:1.24.1-alpine3.21

# Install air
RUN apk add --no-cache git
RUN go install github.com/cosmtrek/air@latest

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 5000
CMD ["air"]