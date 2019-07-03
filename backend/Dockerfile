FROM golang:1.12.5 as builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/znacol/camping/backend

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/backend .

# Second stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /go/bin/backend .

EXPOSE 8000

CMD ["./backend"]