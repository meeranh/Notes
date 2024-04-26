# CMD
+ What `CMD` does is it will run whatever command you specify when the container is started.
+ An important thing to note is that you can only have one `CMD` instruction in a Dockerfile.
+ If you have multiple `CMD` instructions, only the last one will be executed.
```Dockerfile
FROM ubuntu:latest
CMD echo "This will not be executed"
CMD echo "But this will be executed"
```

# CMD Best Practices
+ You can pass in arguments to `CMD` as a normal string, but it is recommended to use the JSON array format.
```Dockerfile
FROM ubuntu:latest
CMD ["echo", "Hello, World!"]
```
+ The above example is the recommended way to use `CMD` and is considered as best practice.

# ENTRYPOINT
+ `ENTRYPOINT` is very similar to `CMD`. It runs commands on container startup.
```Dockerfile
FROM ubuntu:latest
ENTRYPOINT ["echo", "Hello, World!"]
```
+ The above example will print out `Hello, World!`.
+ When run individually, `ENTRYPOINT` and `CMD` do the same thing.

# Using `ENTRYPOINT` & `CMD` Together
+ `CMD` & `ENTRYPOINT` are meant to be used together.
+ Lets use the following example to base our explanation:
```Dockerfile
FROM ubuntu:latest
ENTRYPOINT ["echo"]
CMD ["Hello, World!"]
```
+ The above example will print out `Hello, World!`.
+ Basically, whatever you pass to `CMD` will be passed as an argument to `ENTRYPOINT`.
+ The interesting thing is that you can override `CMD` when running the container.
+ For the above example, if you run the following command:
```bash
docker container run myimage:v1 "Hello, Docker!"
```
+ The output will be `Hello, Docker!` instead of `Hello, World!`.
+ This is because what you pass into the `docker container run` command will override `CMD`.
+ This feature is very useful if we want to pass arguments such as environment variables to the container.
+ There cannot be multiple `ENTRYPOINT` instructions in a Dockerfile.
+ Same like in `CMD`, if you have multiple `ENTRYPOINT` instructions, only the last one will be executed.

# Adding Shell Files to `ENTRYPOINT`
+ If you create a custom shell script and want to use it as the `ENTRYPOINT`, you can do so by using the following example:
```Dockerfile
FROM ubuntu:latest
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
```

+ Our shell script `entrypoint.sh` will look like this:
```bash
#!/bin/bash
echo "This is the entrypoint script"
```
