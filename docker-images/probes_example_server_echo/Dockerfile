FROM golang:1.26 AS build
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0
RUN go build

FROM debian:13-slim
COPY --from=build /build/probes_example_server_echo /
RUN ln -sf /probes_example_server_echo /usr/local/bin/http
CMD ["/probes_example_server_echo", "server"]
EXPOSE 8000
