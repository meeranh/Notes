# Availability In Perspective
+ Most cloud providers promise 99.99% availability.
+ Some cloud providers even promise 99.999% availability.
+ Some cloud providers promise 95% availability.

+ While these may look like small differences, they are not.
+ The following table shows how much downtime each of these promise per month:

| Availability | Downtime Per Month |
|--------------|---------------------|
| 99.999%      | 26.3 seconds        |
| 99.99%       | 4.38 minutes        |
| 95%          | 36 hours            |

+ You would not want to be in a situation where your cloud provider is down for 36 hours.

### Why Are Such Downtimes Even Expected?
+ As a cloud provider, you will be developing new features.
+ When you push these features to production, you need to take the service down for a while.
+ When you take the service down, you are not providing the service, and that is how downtime is calculated.

# How To Combat Downtime?
+ In Azure, there are two ways you can combat downtime:
    1. **[[Regions & Zones | Availability Zones]]**
    2. **Availability Set**

+ Since [[Regions & Zones | availability zones]] were covered in a different section, let's focus on 'Availability Set' here.

# Availability Sets
+ As we mentioned earlier, when a could provider wants to push new features to production, they need to take the service down.
+ Another way you could get downtime is if the hardware that your VM is running on fails or is faulty.
+ Azure provides **fault domains** and **update domains** to combat this.

## Fault Domains
+ These are basically different racks in the same data center.
+ They run on the same power cables, network switches, etc.
+ When specifying an availability set, you can specify the number of fault domains.
+ Basically, you are telling Azure to run your VM on different racks in the same data center.

## Update Domains
+ These are separate machines in the same data center.
+ When a cloud provider pushes new features to production, they take down one update domain at a time.
+ Therefore, if you have five update domains, the updates will be pushed to one update domain at a time.
+ So, your VM will still be available even when the updates are being pushed.

## Differences Between Fault & Update Domains

| Fault Domains                                 | Update Domains                                |
|-----------------------------------------------|-----------------------------------------------|
| Different physical machines                   | Different logical groupings of machines       |
| Meant to protect against hardware failures    | Meant to protect against planned maintenance  |
