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

COPY cmd/cmd .
COPY --from=pyltjie /arrow/build/*.dart.js dist/js/
COPY --from=styler /scissor/dist/css dist/css

COPY assets/html dist/html
COPY assets/fonts dist/fonts
COPY assets/ico dist/ico

EXPOSE 8093

ENTRYPOINT [ "./cmd" ]