FROM golang:1 AS builder
COPY . /src
WORKDIR /src
ENV CGO_ENABLED 0
RUN go get -d ./...
RUN go build -o /assets/out github.com/dcsg/datetime-version-resource/cmd/out
RUN go build -o /assets/in github.com/dcsg/datetime-version-resource/cmd/in
RUN go build -o /assets/check github.com/dcsg/datetime-version-resource/cmd/check
RUN set -e; for pkg in $(go list ./...); do \
		go test -o "/tests/$(basename $pkg).test" -c $pkg; \
	done

FROM alpine:edge AS resource
RUN apk add --update bash ca-certificates tzdata
COPY --from=builder /assets /opt/resource
RUN chmod +x /opt/resource/*

FROM resource AS tests
RUN apk add --update tzdata
COPY --from=builder /tests /tests
RUN set -e; for test in /tests/*.test; do \
		$test -ginkgo.v; \
	done

FROM resource
