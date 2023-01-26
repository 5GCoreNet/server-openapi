FROM golang:1.10 AS build
WORKDIR /go/src
COPY go github.com/5GCoreNet/server-openapi/VAE_DynamicGroup
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi_VAE_DynamicGroup .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapi_VAE_DynamicGroup ./
EXPOSE 8080/tcp
ENTRYPOINT ["github.com/5GCoreNet/server-openapi/VAE_DynamicGroup_VAE_DynamicGroup"]
