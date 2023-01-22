FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o Eecs_EESRegistration .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/Eecs_EESRegistration ./
EXPOSE 8080/tcp
ENTRYPOINT ["./Eecs_EESRegistration"]
