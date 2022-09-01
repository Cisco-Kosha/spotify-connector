FROM docker.io/golang:alpine as builder
RUN apk add git
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o spotify-connector .
FROM docker.io/alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/spotify-connector /app/
WORKDIR /app
CMD ["./spotify-connector"]