FROM golang AS builder

WORKDIR $GOPATH/src/github.com/z0mi3ie/goimgs
COPY . .
RUN go build -o $GOPATH/bin/goimgs

FROM alpine
COPY --from=builder /go/bin/goimgs /go/bin/goimgs
# VOLUME [ "/Users/mr_trashcans/go/src/github.com/z0mi3ie/goimgs/data/www/images/" ]
RUN ls /go/bin/goimgs 
ENTRYPOINT ["/go/bin/goimgs"]
