FROM golang as builder
LABEL stage=builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o dist/api main.go

FROM golang as publish
WORKDIR /app
COPY --from=builder --chown=1001:1001 /build/dist/api /usr/local/bin/api
ENTRYPOINT [ "api" ]
