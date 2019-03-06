FROM alpine:latest

COPY theme .
COPY conf conf
COPY dist dist

COPY assets/html dist/html
COPY assets/fonts dist/fonts
COPY assets/ico dist/ico
COPY assets/img dist/img

EXPOSE 8093

ENTRYPOINT [ "./theme" ]
