FROM golang:1.18-alpine AS builder

RUN mkdir -p $GOPATH/src/github.com/pablodz/sherlockgo
WORKDIR $GOPATH/src/github.com/pablodz/sherlockgo
ENV GO111MODULE=on
COPY . .
RUN go build ./... && go build

FROM alpine

# Dependencies
RUN apk --no-cache add ca-certificates bash croc curl iproute2 iputils

WORKDIR /root/
COPY --from=builder /go/src/github.com/pablodz/sherlockgo .
RUN chmod +x ./sherlockgo
EXPOSE 6969
ENTRYPOINT [ "sh", "./sherlockgo" ]
