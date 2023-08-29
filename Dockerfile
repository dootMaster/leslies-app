FROM golang:1.20.6

WORKDIR /usr/src/leslies-app

COPY . .

RUN go mod tidy

RUN go build -o main .

CMD "./backend/main"