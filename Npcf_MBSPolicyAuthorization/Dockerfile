FROM golang:1.10 AS build
WORKDIR /go/src
COPY go github.com/5GCoreNet/server-openapi/Npcf_MBSPolicyAuthorization
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi_Npcf_MBSPolicyAuthorization .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapi_Npcf_MBSPolicyAuthorization ./
EXPOSE 8080/tcp
ENTRYPOINT ["github.com/5GCoreNet/server-openapi/Npcf_MBSPolicyAuthorization_Npcf_MBSPolicyAuthorization"]
