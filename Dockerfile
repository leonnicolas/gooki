FROM gcr.io/distroless/static

ARG TARGETARCH
COPY ./bin/linux/${TARGETARCH}/gooki /bin/gooki

ENTRYPOINT ["/bin/gooki"]
