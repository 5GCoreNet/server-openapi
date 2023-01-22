FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o M1_EdgeResourcesProvisioning .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/M1_EdgeResourcesProvisioning ./
EXPOSE 8080/tcp
ENTRYPOINT ["./M1_EdgeResourcesProvisioning"]
