# theme
The Template, CSS and JS for Mango modules.

## Run with Docker
*$ GOOS=linux GOARCH=amd64 go build
*$ gulp
*$ docker build -t avosa/theme:dev .
*$ docker rm themeDEV
*$ docker run -d --network host --name themeDEV avosa/theme:dev 
*$ docker logs themeDEV
