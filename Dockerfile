FROM golang:1.17-alpine as builder
WORKDIR /build

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 go build

# Create final image
FROM alpine:latest
WORKDIR /
COPY --from=builder /build/bowditch .
EXPOSE 3000
ENTRYPOINT ["./bowditch"]