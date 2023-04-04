FROM golang:1.19.7-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/tokend/notifications-router-svc
COPY . .

RUN GOOS=linux go build -o /usr/local/bin/notifications-router-svc /go/src/gitlab.com/tokend/notifications-router-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/notifications-router-svc /usr/local/bin/notifications-router-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["notifications-router-svc"]
