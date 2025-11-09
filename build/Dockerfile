# --- build stage ---
FROM golang:1.25.3 AS builder
WORKDIR /src

# Cache deps first
COPY go.mod go.sum ./
RUN go mod download

# Bring in the rest
COPY . .

# Build the examples/colors program as a static binary
RUN go build -o /out/colors ./examples/colors

# --- run stage ---
FROM alpine:latest
COPY --from=builder /out/colors /colors
ENTRYPOINT ["/colors"]
