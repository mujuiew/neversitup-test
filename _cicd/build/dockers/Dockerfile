FROM golang:latest

WORKDIR /app

# COPY go.mod .
# COPY go.sum .
# RUN go mod download
COPY . .

RUN go build -v -mod=vendor -o go/servapi ./cmd/servapi

ENTRYPOINT ["./go/servapi"]
