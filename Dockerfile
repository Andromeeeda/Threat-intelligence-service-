FROM golang:1.26.0-bookworm AS builder 
WORKDIR /app
COPY go.mod go.sum ./ 
RUN go mod download 
COPY . . 
RUN CGO_ENABLED=0 go build -o /app/cmd/threat-service/exe /app/cmd/threat-service

FROM alpine:3.24.3
WORKDIR /app
COPY --from=builder /app/cmd/threat-service/exe /app 
CMD ["/app/exe"]