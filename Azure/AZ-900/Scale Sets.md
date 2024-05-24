# Load Balancer
+ If you have millions and millions of customers visiting your website, but the website is hosted on just one VM, there is a problem.
+ If that VM goes down, your website goes down.
+ Also, a single VM may not be able to handle all the traffic.
+ This is where a load balancer comes in.
![[Pasted image 20240429114945.png]]

+ A load balancer will let you distribute the incoming traffic across multiple VMs.
+ Doing so will ensure that the traffic is evenly distributed across multiple VMs and that no single VM is overwhelmed.

# Scale Sets
+ These are a set of identical VMs.
+ Lets say you created a VM, 
