# Serverupdater
Provides a small webserver that can be used as a target of a github hook.
I am using it to update my server configuration based on the changes.

# Usage
## Update-Docker-Compose
The build-config provides the route `/update` which will trigger a `docker-compose pull` and `docker-compose up -d` on the server.
In order to make it work you need to run the container with the directory where the `docker-compose.yml` is located mounted to `/mount` and the docker-deamon injected like so:
```
docker run -p 80:80 \
    -v /my-docker-compose-repo/:/mount 
    -v /var/run/docker.sock:/var/run/docker.sock 
    serverupdater:latest
```
