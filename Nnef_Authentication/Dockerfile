FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o Nnef_Authentication .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/Nnef_Authentication ./
EXPOSE 8080/tcp
ENTRYPOINT ["./Nnef_Authentication"]
