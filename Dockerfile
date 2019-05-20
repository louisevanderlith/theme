FROM golang:1.11 as build_base

WORKDIR /box

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base as builder

COPY main.go .
COPY controllers ./controllers
COPY core ./core
COPY routers ./routers

RUN CGO_ENABLED="0" go build

FROM alpine:latest AS styler

RUN apk --no-cache add nodejs nodejs-npm
RUN npm install -g gulp gulp-cli

WORKDIR /scissor
COPY package.json .
COPY package-lock.json .
RUN npm install

COPY gulpfile.js .
COPY assets/css ./assets/css
RUN gulp --tasks
RUN gulp

FROM google/dart AS pyltjie
ENV PATH="$PATH:/root/.pub-cache/bin"

WORKDIR /arrow
COPY web ./web
COPY pubspec.yaml pubspec.yaml

RUN pub global activate webdev
RUN pub get
RUN webdev build

FROM scratch

COPY --from=builder /box/theme .
COPY --from=pyltjie /arrow/build/*.dart.js dist/js/
COPY --from=styler /scissor/dist/css dist/css
COPY conf conf

COPY assets/html dist/html
COPY assets/fonts dist/fonts
COPY assets/ico dist/ico

EXPOSE 8093

ENTRYPOINT [ "./theme" ]