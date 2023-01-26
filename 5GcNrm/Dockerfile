FROM golang:1.10 AS build
WORKDIR /go/src
COPY go github.com/5GCoreNet/server-openapi/5GcNrm
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi_5GcNrm .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapi_5GcNrm ./
EXPOSE 8080/tcp
ENTRYPOINT ["github.com/5GCoreNet/server-openapi/5GcNrm_5GcNrm"]
