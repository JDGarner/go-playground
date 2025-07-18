# =========================
# STAGE 1: build go app
# =========================
FROM golang:1.24 AS build

ADD . /app
WORKDIR /app
RUN make dep
RUN make build

# =========================
# STAGE 2: runtime
# =========================
FROM alpine:3.17

WORKDIR /app
COPY --from=build /app/playground /app/playground

CMD ["./playground"]
