# Running Containers
+ After you are done building an image, it is time to run a container.
+ You can use the following command to run a container:
```bash
docker container run <image-name>:<tag>
```

# Running Containers With Port Forwarding
+ You may want to run a container and forward a port from the container to the host machine.
+ An example of this would be to expose your HTTP server running on port 80 to the host machine.
+ You can do this by using the `-p` flag.
+ In the following example, we are forwarding port 80 on the container to port 8080 on the host machine:
```bash
docker container run -p 8080:80 <image-name>:<tag>
```
+ So the standardized syntax for forwarding a port is:
```bash
docker container run -p <host-port>:<container-port> <image-name>
```

# Assigning Random Ports To Containers
+ If you want Docker to assign a random port on the host machine, you can use the `-P` flag.
+ The `-P` flag will publish all exposed ports to random ports on the host machine.
+ An example of this would be:
```bash
docker container run -P <image-name>:<tag>
```

# Running A Container In The Background (Detached Mode)
+ By default, when you run a container, it will run in the foreground.
+ This means that you will see the output of the container in your terminal.
+ If you want to run a container in the background, you can use the `-d` flag.
+ An example of this would be:
```bash
docker container run -d <image-name>:<tag>
```

# Attaching A Detached Container
+ If you have a container running in the background and you want to attach to it, you can use the `attach` command.
+ First, you will have to get the container ID by running `docker container ls`.
+ Then you can attach to the container by running:
```bash
docker container attach <container-id>
```

# Getting Shell Access To A Container
+ If you want to get shell access to a container, you can use the `exec` command.
+ As before, you will have to get the container ID by running `docker container ls`.
+ Then you can get shell access to the container by running:
```bash
docker container exec -it <container-id> /bin/bash
```
+ Lets break down `-it` flag and what it does:
    + `-i`
        + This flag is used to keep STDIN open even if the container is not attached.
        + This is useful for sending input to the container.
    + `-t`
        + This flag is used to allocate a pseudo-TTY.
        + You will be able to have command-line interaction with the container.

