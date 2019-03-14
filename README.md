# theme
The Template, CSS and JS for Mango modules.

## Run with Docker
* $ docker build -t avosa/theme:latest .
* $ docker rm ThemeDEV
* $ docker run -d -e RUNMODE=DEV -p 8093:8093 --network mango_net --name ThemeDEV avosa/theme:latest
* $ docker logs ThemeDEV