# Ensuring Sufficient Availability/Infrastructure Pre-Cloud
+ Assume that you were running a web application for a store during the start of the internet.
+ The only way you can ensure high availability is by purchasing two servers and running them in parallel.
+ Also, if you wanted to anticipate a surge in traffic, you would have to purchase additional servers.
+ This comes with a set of drawbacks:

| Drawback              | Description                                                           |
|-----------------------| ----------------------------------------------------------------------|
| Cost                  | Servers are expensive to purchase and maintain.                       |
| Scalability           | You have to purchase additional servers to scale.                     |
| Availability          | If one server goes down, the other server has to take over.           |
| Low Latency           | If traffic comes from a different country, the latency will be high.  |
| Wasted Infrastructure | You waste money if the servers are not used to their full capacity.   |

# Regions
+ In the above table, *Low Latency* happens because the servers are located in a single location.
+ If traffic comes from America, but the servers are located in Sri Lanka, the latency will be high.
+ Regions in Azure are like locations where you can deploy your resources.
+ Therefore, you can deploy your VMs, databases, etc., in multiple regions to reduce latency.
+ For example, you can deploy your resources to 'East US', 'Australia East', 'Central India', etc.
+ This will ensure that you have a low latency for users in those regions.
+ This also ensures that you meet some compliance requirements (no data should leave the country).

# Availability Zones
+ Availability zones are like regions within a region.
+ Lets say you have deployed your resources in 'East US'.
+ You still are not safe because if the data center in 'East US' goes down, your availability is affected.
+ You now want to deploy your resources in multiple data centers within 'East US'.
+ Availability zones let you do this. It lets you deploy your resources in multiple data centers within a region.
+ Now, you can be sure that if one data center goes down, the other data center will take over.
+ **Note**: Not all regions have availability zones. You have to check if the region you are deploying to has availability zones.
