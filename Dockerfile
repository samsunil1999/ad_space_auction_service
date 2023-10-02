FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/src/app ./...

EXPOSE 8080

CMD ["./ad_space_auction_service"]