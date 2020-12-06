FROM golang:1.13 as builder
WORKDIR /greeter
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build main.go

FROM alpine:latest as prod
COPY --from=builder /greeter/main .
EXPOSE 50051
CMD ["./main"]
