FROM golang:1.25-alpine AS build
WORKDIR /src
COPY main.go .
RUN CGO_ENABLED=0 go build -o /katalog main.go

FROM alpine:3.22
WORKDIR /srv
COPY --from=build /katalog /usr/local/bin/katalog
COPY --chown=10001:10001 data/ ./data/
USER 10001:10001
EXPOSE 8080
CMD ["katalog"]
