FROM golang:1.10 AS build
WORKDIR /go/src
COPY go github.com/5GCoreNet/server-openapi/N5g-eir_EquipmentIdentityCheck
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi_N5g-eir_EquipmentIdentityCheck .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapi_N5g-eir_EquipmentIdentityCheck ./
EXPOSE 8080/tcp
ENTRYPOINT ["github.com/5GCoreNet/server-openapi/N5g-eir_EquipmentIdentityCheck_N5g-eir_EquipmentIdentityCheck"]
