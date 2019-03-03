FROM alpine:latest

COPY theme .
COPY conf conf
COPY dist dist

COPY assets/html dist/html
COPY assets/fonts dist/fonts

EXPOSE 8093

ENTRYPOINT [ "./theme" ]
