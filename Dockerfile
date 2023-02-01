FROM golang:alpine as builder
WORKDIR /app
ADD . .
RUN go build .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app .
EXPOSE 3000
CMD ["./appcast"]
