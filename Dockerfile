FROM golang:alpine AS builder

LABEL maintainer="ankur@opennetworking.org"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

#ENV PORT 6000

RUN go build

FROM alpine

#RUN apk update

WORKDIR /bin

COPY --from=builder /app/upf-adapter /bin/

CMD ["./upf-adapter"]
