FROM golang:1.20-alpine

WORKDIR /app

COPY bankleitzahlen.txt ./
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /kasi-bank-info

CMD [ "/kasi-bank-info" ]