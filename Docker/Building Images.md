# What is A Dockerfile?
+ When you want to create a container, you need to create a Dockerfile.
+ Inside the Dockerfile, you will specify the base image you need and the commands you want to execute.
+ For example, take the following Dockerfile:
```Dockerfile
FROM ubuntu:latest
CMD echo "Hello, World!"
```

+ What this basically does is it tells Docker to use the latest version of Ubuntu as the base image and then run the command `echo "Hello, World!"` when the container is started.

# Building An Image
+ Once you have written your Dockerfile, you can build the Docker image using the following syntax:
```bash
docker image build -t <image-name>:<tag> <path-to-Dockerfile>
```

+ So in our case, we can run the following command and build and run the image:
```bash
docker image build -t myimage:v1 .
docker container run myimage:v1
```

# How Does `RUN` Work?
+ The `RUN` command is used to execute commands inside the Docker container.
+ All `RUN` commands specified in the Dockerfile will be executed **when the image is built**, not when the image is run.
```Dockerfile
FROM ubuntu:latest
RUN apt update
RUN apt install -y curl
RUN apt install -y python3
```

+ When you run `docker image build -t myimage:v1 .`, the `RUN` commands will be executed on build time.
+ So in our case, when the image is being built, it will update the package list, install `curl`, and install `python3`.

# Building Dockerfiles With Custom Names
+ You may have Dockerfiles with names such as `Dockerfile.prod` for `myDockerFile`.
+ If you have custom Dockerfile names, you can use the following syntax to build the image:
```bash
docker image build -f <path-to-Dockerfile> <path-to-context> -t <image-name>:<tag>
```
+ The key difference here is that we are passing a `-f` flag and a `<path-to-Dockerfile>` along with a `<path-to-context>`.
+ What `<path-to-context>` means is the path to the directory where the image should be building from.
+ If you were in a different directory, you would have to specify the path to the directory where the Dockerfile is located.
+ A very simple example is as follows:
```bash
docker image build -f Dockerfile.prod . -t myimage:v1
```

# Building Dockerfiles From Remote Git Repositories
+ If you have a Dockerfile in a Git repository online, you won't have to clone it to build the image.
+ You can build the image using the following syntax:
```bash
docker image build <git-repo-url> -t <image-name>:<tag>
```
