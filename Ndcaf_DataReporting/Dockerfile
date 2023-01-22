FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o Ndcaf_DataReporting .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/Ndcaf_DataReporting ./
EXPOSE 8080/tcp
ENTRYPOINT ["./Ndcaf_DataReporting"]
