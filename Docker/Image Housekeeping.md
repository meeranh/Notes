# Listing All Images
+ To list all the images on your machine, you can run the following command:
```bash
docker image ls
```
+ You can list all the images on your machine with the `-a` flag:
```bash
docker image ls -a
```

# Deleting An Image
+ To delete an image, you can use the `image rm` command followed by the image ID or image name.
+ For example, if you want to delete an image with the ID `1234567890ab`, you would run:
```bash
docker image rm 1234567890ab
```
+ A shorthand for the `image rm` command is `rmi`.
+ So you can also run:
```bash
docker image rmi 1234567890ab
```
+ You do not always need to delete an image by its ID. You can also delete an image by its name.
+ For example, if you want to delete an image called `myimage`, you would run:
```bash
docker image rm myimage
```

# Deleting Images Forcefully
+ If you have a container running from an image, you cannot delete the image.
+ You will have to stop the container first and then delete the image.
+ If you want to forcefully delete an image, you can use the `-f` flag.
+ What this will do is Docker will stop all containers running from the image and then delete the image.
+ For example, if you want to forcefully delete an image with the ID `1234567890ab`:
```bash
docker image rm -f 1234567890ab
```

# Deleting All Images
+ If you want to delete all images on your machine, you can use the following command:
```bash
docker image rm $(docker image ls -aq)
```

+ If some images are being used by containers and you want to forcefully delete them, you can use the `-f` flag:
```bash
docker image rm $(docker image ls -aq) -f
```

# Removing All Unused Images (Pruning Dangling Images)
+ If you want to remove all unused images on your machine, you can use the `image prune` command.
+ The `image prune` command will remove all dangling images (images that are not being used by any container).
+ To remove all dangling images, you can run:
```bash
docker image prune
```
