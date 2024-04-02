FROM golang:1.22-alpine as builder

# build
ARG TARGETOS
ARG TARGETARCH

# Add libraries
RUN apk add --no-cache git make ca-certificates

# Setup
WORKDIR /build

# download modules
COPY go.mod .
COPY go.sum .
RUN go mod download


# Copy & build
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} GO111MODULE=on make build

# Copy into scratch container
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/traefik-forward-auth ./
ENTRYPOINT ["./traefik-forward-auth"]
