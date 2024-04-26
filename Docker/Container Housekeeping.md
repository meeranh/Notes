# View All Running Containers
+ To view all the running containers on your machine, you can run the following command:
```bash
docker container ls
```
+ The same can be achieved with the `ps` command:
```bash
docker container ps
```

# Stopping A Container
+ To stop a container, you can use the `stop` command followed by the container ID or container name.
+ For example, if you want to stop a container with the ID `1234567890ab`, you would run:
```bash
docker container stop 1234567890ab
```

# Removing A Container
+ To delete a container, you can use the `rm` command followed by the container ID or container name.
+ For example, if you want to delete a container with the ID `1234567890ab`, you would run:
```bash
docker container rm 1234567890ab
```

+ If you force the deletion of a container, you can use the `-f` flag.
+ For example, if you want to forcefully delete a container with the ID `1234567890ab`:
```bash
docker container rm -f 1234567890ab
```

# Removing All Containers
+ If you want to delete all containers on your machine, you can use the following command:
```bash
docker container rm $(docker container ls -aq)
```

# Remove All Containers That Are Not Running (Pruning Stopped Containers)
+ If you want to remove all containers that are not running, you can use the `container prune` command.
+ The `container prune` command will remove all stopped containers.
+ To remove all stopped containers, you can run:
```bash
docker container prune
```
