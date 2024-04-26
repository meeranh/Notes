# History of Application Deployment
+ At the start of the 21st century, applications were deployed on physical servers.
+ This comes with a new set of problems:
    1. **Resource Utilization**: Physical servers are expensive and underutilized.
    2. **Scalability**: It is difficult to scale physical servers.
    3. **Isolation**: Applications running on the same server can interfere with each other.
    4. **Dependency Management**: Applications can have conflicting dependencies.
    5. **Portability**: Applications are not portable across different environments.
+ To solve these problems, virtualization was introduced.

# The Problem With Virtualization
+ Virtualization is when you run multiple virtual machines on a single physical server.
+ This will let you create different operating system instances.
+ Because we now have different operating system environments, we solve the problem of dependency management.
+ We can now also isolate applications from each other.
+ Finally, the issue with scalability and resource allocation is gone because:
    1. We can now run multiple virtual machines on a single physical server.
    2. We can upscale/downscale the RAM and CPU of each virtual machine.

+ This introduced another problem:
    1. **Resource Overhead**: Each virtual machine has its own operating system, which consumes resources.
    2. **Slow Boot Time**: Booting a virtual machine can take minutes.
    3. **Large Disk Space**: Each virtual machine has its own disk space.

+ In short, because each VM has its own operating system, it consumes a lot of resources.
+ To put this in perspective, imagine you have a physical machine running Ubuntu and it has 8GB of RAM.
+ We also have 2 VMs running Ubuntu, each with 2GB of RAM.
+ Assuming that Ubuntu consumes 1GB of RAM:
    + The first VM will waste 1GB of RAM for its operating system.
    + The second VM will waste 1GB of RAM for its operating system.
    + The physical machine will waste 1GB of RAM for its operating system.
+ In summary, out of the 8GB, 3GB is wasted on operating systems.

# How Docker Works
+ Docker tries to solve the issue we have with virtualization.
+ It knows that Ubuntu, Debian, CentOS, and other distributions are all Linux distributions.
+ In theory, they should technically use the same Linux kernel.
+ Docker understands this and assigns the same kernel as the host machine to all the applications it is running.
+ Doing so will eliminate the need for each application to have its own operating system.
+ If we compare virtualization vs Docker, we now save 3GB of RAM, which could be used for other applications.

# Image
+ As said before, we won't be installing an entire operating system (including its kernel) on each application.
+ We will only install the application and its dependencies.
+ Lets say that we wanted to install Ubuntu, its package manager is `apt`.
+ Docker will fetch the Ubuntu image from the Docker Hub.
+ All this image contains is the basic packages for Ubuntu and the `apt` package manager.
+ Therefore, it is lightweight by nature.

# Container
+ Once an image is downloaded from places like Docker Hub, it is stored on the host machine.
+ These containers can be run on any machine that has Docker installed.
+ What this does is it assigns the same kernel as the host machine to the container.
+ The container will run the packages and dependencies that are in the image over the host machine's kernel.
