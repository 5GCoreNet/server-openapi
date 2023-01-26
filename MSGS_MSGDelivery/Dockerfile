FROM golang:1.10 AS build
WORKDIR /go/src
COPY go github.com/5GCoreNet/server-openapi/MSGS_MSGDelivery
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi_MSGS_MSGDelivery .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapi_MSGS_MSGDelivery ./
EXPOSE 8080/tcp
ENTRYPOINT ["github.com/5GCoreNet/server-openapi/MSGS_MSGDelivery_MSGS_MSGDelivery"]
