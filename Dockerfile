from golang:1.21 as builder

WORKDIR /src

COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN make build

FROM gcr.io/distroless/static

COPY --from=builder /src/gooki /bin

ENTRYPOINT ["/bin/gooki"]
