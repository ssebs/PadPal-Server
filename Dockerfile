FROM golang:latest

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./ ./

RUN go build -o ./bin/ ./...
RUN chmod a+x bin/cmd

EXPOSE 80

CMD ["/app/bin/cmd"]