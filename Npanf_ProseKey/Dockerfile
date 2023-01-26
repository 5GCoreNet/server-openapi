FROM golang:1.10 AS build
WORKDIR /go/src
COPY go github.com/5GCoreNet/server-openapi/Npanf_ProseKey
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi_Npanf_ProseKey .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapi_Npanf_ProseKey ./
EXPOSE 8080/tcp
ENTRYPOINT ["github.com/5GCoreNet/server-openapi/Npanf_ProseKey_Npanf_ProseKey"]
