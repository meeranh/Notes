# Building Named Containers
+ When you build a Docker image, you can assign a name to the container.
+ This is useful when you want to reference the container by name instead of the container ID.
+ The `--name` flag is used to assign a name to the container.
+ The syntax for building a named container is:
```bash
docker container run --name <container-name> <image-name>:<tag>
```
