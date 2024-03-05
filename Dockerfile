FROM eu.gcr.io/connect-f7e5b/elective/golang:alpine as builder

WORKDIR /graph-ql-api

# install make and git command
RUN apk add --virtual --no-cache git curl build-base

# add this entire directory (excl .dockerignore content) to the working directory
ADD . .

# build the application
RUN make build

EXPOSE 80

CMD ["/graph-ql-api/graph-ql-api"]
# Stage 2: Potentially brittle, as we go on the assumption that the golang:alpine
#  container will always have a $GOPATH of /go/. However, we can reduce the total
#  image size to <9mb, and the $GOPATH of the golang:alpine image seems consistent
#  between versions.
FROM eu.gcr.io/connect-f7e5b/elective/alpine:latest

WORKDIR /
COPY --from=builder /graph-ql-api/graph-ql-api /graph-ql-api

CMD ["/graph-ql-api"]
