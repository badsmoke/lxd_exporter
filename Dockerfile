FROM golang:1.15-alpine AS builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN apk add make git gcc libc-dev 
WORKDIR /src
RUN mkdir bin
COPY . .

RUN go build -o bin/ ./...

FROM alpine:latest
RUN apk --no-cache add ca-certificates


#permission denied socket
#RUN addgroup -g 1001 appgroup && addgroup -g 131 lxd && \
#  adduser -H -D -s /bin/false -G appgroup -G lxd -u 1001 appuser
#USER 1001:1001

COPY --from=builder /src/bin/lxd_exporter /bin/lxd_exporter
ENTRYPOINT ["/bin/lxd_exporter"]