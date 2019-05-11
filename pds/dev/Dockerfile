# Build the project
FROM golang:1.12.0 as builder

WORKDIR /go/src/github.com/dbmedialab/truesight-api
ADD . .

RUN make build
RUN make test

# Create production image for application with needed files
FROM golang:1.12.0-alpine3.9

EXPOSE 8080

RUN apk add --no-cache ca-certificates bash

COPY --from=builder /go/src/github.com/dbmedialab/truesight-api/main .

CMD ["./main"]
