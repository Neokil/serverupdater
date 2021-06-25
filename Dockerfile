#
# Build the Application
#
FROM golang:1.16-alpine AS build

RUN apk add --update gcc musl-dev

WORKDIR /go/src
COPY . .

# The "normal" command does not work since sqlite needs CGO
# RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/server cmd/server/server.go
RUN go build -a -ldflags '-linkmode external -extldflags "-static"' -o build/server cmd/server/server.go

WORKDIR /go/src/build
CMD [ "server"]
EXPOSE 80

#
# Create a Release-Image from Emptry-Scratch
#
FROM docker:dind as release
COPY --from=build /go/src/build /
RUN apk add docker-compose
WORKDIR /mount
ENTRYPOINT [ "/server", "--config=/config.json"]
EXPOSE 80