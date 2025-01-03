FROM golang:1.23 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
ARG BUILD_TARGET
RUN go build -v -o app.bin ${BUILD_TARGET}

FROM debian:12 AS final
ENV TZ=UTC
WORKDIR /app
COPY --from=builder /src/app.bin /app/app.bin
CMD ["/app/app.bin"]
