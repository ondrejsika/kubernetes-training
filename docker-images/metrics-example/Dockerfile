FROM golang as build
WORKDIR /build
COPY go.mod .
COPY go.sum .
COPY server.go .
RUN go get
ENV CGO_ENABLED=0
RUN go build server.go

FROM scratch
COPY --from=build /build/server .
CMD ["/server"]
EXPOSE 80
