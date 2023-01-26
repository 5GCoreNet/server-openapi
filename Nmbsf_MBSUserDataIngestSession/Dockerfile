FROM golang:1.10 AS build
WORKDIR /go/src
COPY go github.com/5GCoreNet/server-openapi/Nmbsf_MBSUserDataIngestSession
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi_Nmbsf_MBSUserDataIngestSession .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapi_Nmbsf_MBSUserDataIngestSession ./
EXPOSE 8080/tcp
ENTRYPOINT ["github.com/5GCoreNet/server-openapi/Nmbsf_MBSUserDataIngestSession_Nmbsf_MBSUserDataIngestSession"]
