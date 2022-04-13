FROM golang:alpine
RUN apk add git

ENV GOOS=linux \
    GOARCH=amd64

WORKDIR /

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o main .
CMD ["/main"]