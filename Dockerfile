FROM golang:alpine AS builder

# multiple stage practice

WORKDIR /go/src/app
COPY . .

RUN go build -o crawler-video .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/ /app/

CMD ["./crawler-video"]