FROM golang:1.10 AS build
WORKDIR /go/src
COPY go github.com/5GCoreNet/server-openapi/Naf_ProSe
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi_Naf_ProSe .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapi_Naf_ProSe ./
EXPOSE 8080/tcp
ENTRYPOINT ["github.com/5GCoreNet/server-openapi/Naf_ProSe_Naf_ProSe"]
