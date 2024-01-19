FROM golang:latest as build
WORKDIR $GOPATH/src/github.com/ca-gip/kubi-members-v2
COPY . $GOPATH/src/github.com/ca-gip/kubi-members-v2
RUN make build

FROM scratch
WORKDIR /root/
COPY --from=build /go/src/github.com/ca-gip/kubi-members-v2/build/kubi-members-v2 .
EXPOSE 8000
CMD ["./kubi-members"]