# multistage build to shrink docker image size
FROM golang:alpine as build

# to support http get requests
RUN apk --no-cache add ca-certificates

WORKDIR /build 

COPY . .

# create binary but remove debug information from binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags '-static'" -o ./app

# use upx to compress binary
RUN apk --no-cache add upx
RUN upx ./app


FROM scratch

LABEL org.opencontainers.image.source https://github.com/cs50victor/scrapperAPI

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/app /app

ENV PORT 3000
EXPOSE 3000

ENTRYPOINT ["/app"]
