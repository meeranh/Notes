# Docker Repository / Container Registry
+ What these are basically is a repository (same like a Git repository) for Docker images.
+ If you built a Docker image, and you want to share it with others, you can push it to a Docker repository.
+ The most popular Docker repository is [Docker Hub](https://hub.docker.com/).

# How to Upload To Docker Hub
+ First, you need to create an account on Docker Hub.
+ Create a repository on Docker Hub.
+ Use the `docker login` command to login to your account on your terminal.
+ Then upload the repository:
    1. Tag the image with your Docker Hub username and the repository name:
        + `docker tag <image_id> <docker_hub_username>/<repository_name>:<tag>`
    2. Push the image to the repository:
        + `docker push <docker_hub_username>/<repository_name>:<tag>`
