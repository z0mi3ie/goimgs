# Article on this approach  https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/z0mi3ie/goimgs
RUN pwd
COPY . .
RUN go get -d -v
RUN go build -o $GOPATH/bin/goimgs

FROM alpine
COPY --from=builder /go/bin/goimgs /go/bin/goimgs
# VOLUME [ "/Users/mr_trashcans/go/src/github.com/z0mi3ie/goimgs/data/www/images/" ]
RUN ls /go/bin/goimgs 
ENTRYPOINT ["/go/bin/goimgs"]