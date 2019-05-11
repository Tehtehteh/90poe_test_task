# Build the project
FROM golang:1.12.0 as builder

WORKDIR /go/src/github.com/tehtehteh/t_task
ADD . .

RUN make build
RUN make test

# Create production image for application with needed files
FROM golang:1.12.0-alpine3.9

EXPOSE 8080

RUN apk add --no-cache ca-certificates bash

COPY --from=builder /go/src/t_task/main .

CMD ["./main"]
