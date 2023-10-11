---
id: "Section 11 - Network Design"
aliases:
  - "The OSI Model"
tags: []
---

# The OSI Model
+ These are the components of the OSI model from bottom-to-top:
    1. **Physical**: Network cables/radio waves transferring data using **bits**.
    2. **Data Link**: Describes how connection is established over a physical layer using **frames**.
    3. **Network**: Routing and switching information such as subnetting is decided here using **packets**.
    4. **Transport**: Method in which data is transferred using TCP/UDP.
    5. **Session**: Handles the connection instance (establishes or terminates a connection).
    6. **Presentation**: Converts data to a readable format for the application layer.
    7. **Application**: The layer where the data was originally created (HTTP, SMTP, etc).

# Switches (Layer 2)
+ A combination of Hubs an Bridges which will only send data to the intended recipient over a large network.
+ Some of the attack on switches are as follows:
    1. **MAC Flooding**: When a hacker floods the switch with fake MAC addresses.
    2. **MAC Spoofing**: When a hacker changes their MAC address to bypass MAC filtering.
        + Limit maximum static MAC addresses
        + Limit duration time of MAC addresses
        + Conduct ARP inspection
    3. **Physical Tampering**: When a hacker physically tampers with the switch.

# Router (Layer 3)
+ Routers decide the route a packet should take from source to destination.
+ **IP Spoofing** is an attack that can bypass access control lists in routers by setting a fake IP address in the packet.

# Network Zones
+ **De-Militarized Zone (DMZs)** are used to store public-facing servers and are not trusted by the internal network.
**Sub-Zones** are used to separate different departments in a company's internal network.
**Extranet** is a network that only a company's partners can access.
**Intranet** is a network that only a company's internal employees can access.

# Jumpbox
+  The hosts in the DMZ should be **Bastion Hosts**, which are servers in the DMZ that are not configured with any services that run on the local network.
+ A **Jumpbox** is a hardened server that connects the internal hosts with the hosts in the DMZ for communication.

# Network Access Control
 + A scan which checks if a device is compliant with the company's security policies before allowing it to connect to the network.
 + There are two types of agents which can connect to a network:
    1. **Persistent Agents**: A software installed on the user's device requesting access to the network.
    2. **Non-persistent Agents**: Software installed on a server or is remotely installed on the device and removed after the scan.
+ **IEEE 802.1x** is a port-based NAC.

# VLANs
+ These are segments of the network that differentiate communications between company departments.
+ Some common atacks against VLANs are:
    1. **Switch Spoofing**: Attacker pretends to be a switch and reroutes traffic maliciously
    2. **Double Tagging**: Network data has two tags, the first tag is removed after processing, the second tag is used to bypass security controls.

# Subnetting
+ This is the logical separation of a network into smaller networks.
+ **Separate policies for separate subnets** can be assigned.
+ If an attack occurs, it only affetcts that specific subnet with **reduced attack surface**.

# Network Address Translation (NAT)
+ Process of **changing IP addresses to another IP address** when a data packet is transferred over multiple routers.
+ It **hides our IP address**.
+ The following are the IP address classes:

| Class |               Range               | 
| :---: | :-------------------------------: |
| **A** | *10.0.0.0* -> *10.255.255.255*    |
| **B** | *172.16.0.0* -> *172.31.255.255*  |
| **C** | *192.168.0.0* -> *192.168.255.255*|

# Telephony
+ This is the transfer of voice data to users.
+ **Modems** were used to convert analog signals to digital signals and were a great attack vector for war dialing.
    + **Dial-up resources** are servers that use telephony to manage the network (call and control the server).
    + **War Dialing** is when you try all possibile phone numbers to access a dial-up resource. Prevention for this is to stop using modems and rely on something like SSH.

+ **Private Branch Exchange (PBX)** runs all the internal phone lines for a company.
    + If attackers exploit PBX systems, they will be able to make calls for free on the company's account.
    + To secure this:
        + Prevent physical access to unauthorized users in PBX server rooms.
        + Remove remote access to PBX servers.
        + Replace PBX with VoIP.
        + Improve good QoS (Quality of Service) to prevent DoS attacks.
