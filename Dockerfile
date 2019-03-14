# Followed https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
# since I was unfamiliar with multistage docker builds, the article continues with optimizations and 
# and security but this will get our app dockerized up without a 300MB container XD 

FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/z0mi3ie/goimgs
RUN pwd
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v

# Build the binary.
RUN go build -o $GOPATH/bin/goimgs

############################
# STEP 2 build a small image
############################
FROM alpine
COPY --from=builder /go/bin/goimgs /go/bin/goimgs
# This does not share to the host
# VOLUME [ "/Users/mr_trashcans/go/src/github.com/z0mi3ie/goimgs/data/www/images/" ]
RUN ls /go/bin/goimgs 
ENTRYPOINT ["/go/bin/goimgs"]