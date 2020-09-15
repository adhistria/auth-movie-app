FROM golang:1.14 as builder

RUN mkdir -p /app 
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o auth-service cmd/auth-movie-app/main.go

FROM alpine
RUN mkdir -p /app
COPY --from=builder /app/auth-service /app

ENTRYPOINT /app/auth-service
EXPOSE 8080