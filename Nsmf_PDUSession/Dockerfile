FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o Nsmf_PDUSession .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/Nsmf_PDUSession ./
EXPOSE 8080/tcp
ENTRYPOINT ["./Nsmf_PDUSession"]
