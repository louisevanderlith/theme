FROM golang:1.13 as build_base

WORKDIR /box

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base as builder

COPY main.go .
COPY handles ./handles
COPY core ./core

RUN CGO_ENABLED="0" go build

FROM alpine:3.12.0 AS styler

RUN apk --no-cache add nodejs nodejs-npm
RUN npm install -g gulp gulp-cli

WORKDIR /scissor
COPY package.json .
COPY package-lock.json .
RUN npm install

COPY assets/css ./assets/css

COPY gulpfile.js .
RUN gulp

FROM google/dart:latest AS pyltjie
ENV PATH="$PATH:/root/.pub-cache/bin"

WORKDIR /arrow
RUN pub global activate webdev

COPY build.yaml build.yaml
COPY pubspec.yaml pubspec.yaml
RUN pub get

COPY web ./web
RUN webdev build

FROM scratch

COPY --from=builder /box/theme .
COPY --from=pyltjie /arrow/build/*.dart.js dist/js/
COPY --from=styler /scissor/dist/css dist/css

COPY assets/colour dist/css
COPY assets/html dist/html
COPY assets/fonts dist/fonts
COPY assets/ico dist/ico

EXPOSE 8093

ENTRYPOINT [ "./theme" ]