FROM golang:1.11 AS builder

WORKDIR /box
COPY go.mod .
COPY go.sum .
RUN go mod download

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
COPY .babelrc .
COPY assets/css ./assets/css

RUN gulp

FROM google/dart AS pyltjie

WORKDIR /arrow
COPY assets/dart ./assets/dart

RUN mkdir -p assets/js
COPY compiledart.sh .
RUN sh ./compiledart.sh

FROM alpine:latest

COPY --from=builder /box/theme .
COPY --from=pyltjie /arrow/assets/js dist/js
COPY --from=styler /scissor/dist/css dist/css
COPY conf conf

COPY assets/html dist/html
COPY assets/fonts dist/fonts
COPY assets/ico dist/ico
COPY assets/img dist/img

EXPOSE 8093

ENTRYPOINT [ "./theme" ]