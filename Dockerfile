FROM golang:alpine as build

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /go/src/go.lafronz.com/vanago
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o runtime ./cmd/vanago/main.go

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/go.lafronz.com/vanago/runtime /app
ENTRYPOINT ["/app"]