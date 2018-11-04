FROM golang:1.8 as build

WORKDIR /go/src/acme-helper
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM scratch

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /go/bin/acme-helper /
CMD ["/acme-helper", '/well-known']
