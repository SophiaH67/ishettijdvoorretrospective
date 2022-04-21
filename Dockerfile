FROM golang:1.16-alpine as builder
WORKDIR /app
COPY ./ ./
ENV CGO_ENABLED=0
RUN go build -o main main.go 

FROM scratch
COPY --from=builder /app/main ./
COPY ./images/ ./images/
ENTRYPOINT ["/main"]
