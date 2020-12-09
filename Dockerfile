FROM golang:1.15-alpine as builder

ARG APP_BUILD_VERSION=""
ARG APP_BUILD_REVISION=""

ENV GOOS=linux \
    GARCH=amd64 \
    CGO_ENABLED=0 \
    GO111MODULE=on

WORKDIR /workspace

COPY . .

RUN apk update && \
    apk add --update --no-cache ca-certificates && \
    go mod download && \
    go mod verify && \
    go build -x -v -a -ldflags "-X main.version=${APP_BUILD_VERSION} -X main.revision=${APP_BUILD_REVISION}" -o /build/app .


FROM scratch

COPY --from=builder /build/app /app

EXPOSE 8000
ENTRYPOINT ["/app"]