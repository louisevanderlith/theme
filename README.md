# theme
The Template, CSS and JS for Mango modules.

## Run with Docker
* $ GOOS=linux GOARCH=amd64 go build
* $ gulp
* $ docker build -t avosa/theme:latest .
* $ docker rm themeDEV
* $ docker run -d -e RUNMODE=DEV -p 8093:8093 --network mango_net --name ThemeDEV avosa/theme:latest
* $ docker logs themeDEV
