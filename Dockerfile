FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/OctaneAL/Shortly
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/Shortly /go/src/github.com/OctaneAL/Shortly


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/Shortly /usr/local/bin/Shortly
RUN apk add --no-cache ca-certificates

EXPOSE 8080

ENTRYPOINT ["sh", "-c", "Shortly migrate up && Shortly run service"]
