FROM golang:1.26 AS builder
ARG BUILD_ID=0
WORKDIR /build
COPY go.mod go.sum ./
RUN \
  --mount=type=cache,target=/go \
  go mod download
COPY . .
RUN \
  --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go \
  CGO_ENABLED=0 go build -ldflags="-X main.BUILD_ID=${BUILD_ID}"

FROM debian:13-slim
COPY --from=builder /build/counter_backend /usr/local/bin/counter_backend
CMD [ "/usr/local/bin/counter_backend" ]
EXPOSE 8000
