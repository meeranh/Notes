# What Are Volume Mounts
+ Lets say that you are working on a project.
+ You have a set of files that you want to copy into your container.
+ You also want to make changes to these files.
+ You can copy these files into your container using the `COPY` command in your Dockerfile.
+ But what if you want to make changes to these files and have them reflected in the container?
+ This is where volume mounts come in.
+ Volume mounts allow you to mount a directory from your host machine into your container.
+ This means that any changes you make to the files in the directory on your host machine will be reflected in the container.

# Docker Volumes
+ To see all your Docker volumes, you can run the following command:
```bash
docker volume ls
```
+ You can find the filepath of a Docker volume using the following command:
```bash
docker volume inspect <volume-name>
```
+ After you have found the filepath of the Docker volume, you can `cd` into the directory and see the contents of the volume.

# Host Volumes
+ Lets say you have a directory called `myfiles` on your host machine.
+ You want to mount this directory into your container.
+ You can use the `-v` flag to mount a directory from your host machine into your container.
+ The syntax for mounting a directory is as follows:
```bash
docker container run -v <host-directory>:<container-directory> <image-name>
```
+ So in our case, we can run the following command to mount the `myfiles` directory into our container:
```bash
docker container run -v /path/to/myfiles:/app <image-name>
```
+ Note that you have to provide the absolute path to the directory on your host machine.

# Anonymous Volumes
+ Lets say you want to build an image and you also want to have its `/home` directory mounted in your system.
+ You can do this with the `-v` flag.
+ The syntax for mounting a directory is as follows:
```bash
docker container run -v <container-directory> <image-name>
```
+ So in our case, we can run the following command to mount the `/home` directory of the container into our system:
```bash
docker container run -v /home <image-name>
```
+ What this will do is it will create a Docker volume and assign the `/home` folder of the container to your Docker volume.
+ You can achieve the same result with a Dockerfile with the `VOLUME` keyword:
```Dockerfile
FROM ubuntu:latest
VOLUME /home
```
+ This is going to create an anonymous volume and assign the `/home` folder of the container to the volume.

# Named Volumes
+ Unlike anonymous volumes, named volumes have a name.
+ The syntax for creating a named volume is as follows:
```bash
docker container run -v <volume-name>:<container-directory> <image-name>
```
+ So in our case, we can run the following command to create a named volume called `myvolume` and assign the `/home` folder of the container to the volume:
```bash
docker container run -v myvolume:/home <image-name>
```
