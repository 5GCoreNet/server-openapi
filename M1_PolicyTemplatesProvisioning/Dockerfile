FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o M1_PolicyTemplatesProvisioning .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/M1_PolicyTemplatesProvisioning ./
EXPOSE 8080/tcp
ENTRYPOINT ["./M1_PolicyTemplatesProvisioning"]