# Copying Files Into The Container
+ Lets say you have a set of files inside a folder.
+ You also want to copy this into your container.
+ You need to use the `COPY` command in your Dockerfile:
```Dockerfile
FROM ubuntu:latest
COPY . /app
```
+ The `COPY` command takes in two arguments:
    1. Source destination on the host machine
    2. Target destination on the Docker container
+ Shortly, the syntax is `COPY <source> <destination>`
+ Providing `COPY . /app` will copy all the files from the current directory on the host machine to the `/app` directory inside the container.

# `COPY` vs `ADD`
+ Both `COPY` and `ADD` do the same thing, which is copy files, but there is a key difference.
+ `COPY` does not extract files from a URL or tarball.
+ `ADD` can do this.
+ For example, lets say that we had a tarball file called `myfiles.tar.gz`:
```Dockerfile
FROM ubuntu:latest
ADD myfiles.tar.gz /app
```
+ What this is going to do is extract the contents of `myfiles.tar.gz` into the `/app` directory inside the container.
+ So you won't have to manually extract the tarball file.
+ But the developers of Docker recommend using `COPY` over `ADD` to download files from the internet.
+ Also, `curl` is a better option for downloading files from the internet such as tarballs.
