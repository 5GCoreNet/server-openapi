FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o 5GcNrm .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/5GcNrm ./
EXPOSE 8080/tcp
ENTRYPOINT ["./5GcNrm"]