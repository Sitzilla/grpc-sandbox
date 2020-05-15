FROM golang:1.13-alpine AS builder
WORKDIR /app

RUN apk add ca-certificates curl protobuf git
RUN curl -sL https://taskfile.dev/install.sh | sh

RUN go get github.com/twitchtv/twirp/protoc-gen-twirp
RUN go get github.com/golang/protobuf/protoc-gen-go

COPY . .

RUN ./bin/task build

FROM scratch

COPY --from=builder /app/bin /bin
COPY --from=builder /etc/ssl /etc/ssl

CMD ["/bin/server"]