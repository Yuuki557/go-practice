FROM golang:1.15.1 as builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY ./package/ ./package/

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
	-o /go/bin/main \
	-ldflags '-s -w' \
	./package

FROM scratch as runner

COPY --from=builder /go/bin/main /app/main

ENTRYPOINT ["/app/main"]