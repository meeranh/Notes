# Port Forwarding
+ Port forwarding is simply the process of forwarding a network port from one network node to another.
+ If we have a Docker container running a service on port 8080, we can forward the port to the host machine.
+ The syntax for port forwarding is `-p <host-port>:<container-port>`.
+ For example, if you want to run the container `myweb:1` and expose port 8080 on the host machine to port 80 on the container, you would run:
```bash
docker container run -p 8080:80 myweb:1
```

+ In short, the syntax for port forwarding is:
```bash
docker container run -p <host-port>:<container-port> <image-name>
```

# `EXPOSE` In Dockerfiles
+ The `EXPOSE` instruction in a Dockerfile is used to inform Docker that the container is listening on a specific port.
```Dockerfile
EXPOSE 80 443
```
+ The above instruction tells Docker that the container listens on ports 80 and 443.
+ However, this does not actually publish the port. To publish the port, you must use the `-p` flag when running the container.
```bash
docker container run -p 8080:80 -p 443:443 myweb:1
```
+ The above command runs the container `myweb:1` and publishes ports 80 and 443 to the host machine.

# Assigning Random Ports
+ If you want Docker to assign a random port on the host machine, you can use the `-P` flag.
+ The `-P` flag will publish all exposed ports to random ports on the host machine.
```bash
docker container run -P myweb:1
```
