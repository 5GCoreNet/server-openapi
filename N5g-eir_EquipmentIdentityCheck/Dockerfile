FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o N5g-eir_EquipmentIdentityCheck .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/N5g-eir_EquipmentIdentityCheck ./
EXPOSE 8080/tcp
ENTRYPOINT ["./N5g-eir_EquipmentIdentityCheck"]
