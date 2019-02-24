# theme
The Template, CSS and JS for Mango modules.

## Run with Docker (linux)
*$ go build
*$ gulp
*$ docker build -t avosa/theme:dev .
*$ docker rm themeDEV
*$ docker run -d --network host --name themeDEV avosa/theme:dev 
*$ docker logs themeDEV

#TODO
setup subdomain and certs
deploy theme.avosa
update dockerfiles of UI's to download from theme.avosa