FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o Ndccf_ContextManagement .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/Ndccf_ContextManagement ./
EXPOSE 8080/tcp
ENTRYPOINT ["./Ndccf_ContextManagement"]
