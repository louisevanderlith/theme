FROM alpine:latest

COPY theme .
COPY conf conf
COPY dist dist

ENTRYPOINT [ "./theme" ]
