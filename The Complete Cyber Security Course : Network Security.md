---
id: The Complete Cyber Security Course : Network Security
aliases:
  - Section 1 & Section 2
tags: []
---

# Section 1 & Section 2

#### Canary Tokens

+ These are URLs which will trigger an action when visited.
+ Some examples of canary tokens are:
    1. [StationX Canary Tokens](https://www.stationx.net/canarytokens/)
    2. [CanaryTokens.org](https://canarytokens.org/generate)

# Section 3: Router Security

+ **Terms**:
    1. **Port**: A doorway into a service of a device.
    2. **Default Gateway**: The internal IP address of the router.
    3. **Network Address Translation**: Translates internal IP addresses to external IP addresses.
    4. **Port Forwarding**: Rerouting traffic from a port on the router to a port on a device.
    5. **MAC Table**: A table of MAC addresses and their corresponding switch port numbers.
    6. **iptables**: A firewall for Linux systems
    7. **Modem**: Converts digital signals to analog signals and vice versa.
    8. **Router**: A device that sends data between networks and its devices.
    9. **Wireless Access Point**: A device that allows devices to connect to a network wirelessly.

+ Finding the default gateway of a router:
    + **Windows**: ```route print```
    + **Linux**: ```sudo route -n```

+ How Network Address Translation (NAT) and Port Forwarding works:
    + **Network Address Translation**: The process of converting internal IPs to a router's external IP.
        1. A device sends data to the router to access the internet.
        2. The router changes the source IP address of the data to its external IP address.
        3. The router sends the data to the internet.

    + **Port Forwarding**: The process of rerouting traffic from a port on the router to a port on a device.
        1. A device sends a packet to the router to access the internet from a specific port.
        2. The router changes the source port of the data to a random port and keeps track of it.
        3. The router sends the data to the internet.
        4. The router receives a response from the internet.
        5. The router changes the destination port of the response to the port of the device.
        6. This is then sent to the device.

+ How switches route data using the MAC table:
    + Imagine a MAC table has the following entries
        1. MAC Address "00:00:00:00:00:31" assigned to switch port 1 
        2. MAC Address "00:00:00:00:00:29" assigned to switch port 2
        3. MAC Address "00:00:00:00:00:03" assigned to switch port 3
        4. MAC Address "00:00:00:00:00:93" assigned to switch port 4

    + If a packet with destination MAC address ```00:00:00:00:00:03``` is received, it will be sent to port 3.
    + This ensures that the switch does not have to broadcast the packet to all ports.

+ **Demilitarized Zone (DMZs)**
    + This is a section of the network which can be accessed by the public internet.
    + It is not recommended to place routers, switches, and sensitive servers in the DMZ.
    + This is because search engines such as Shodan will find these devices.

+ **Universal Plug and Play (UPnP)**
    + This is a feature that automates port forwarding.
    + If we have a device in a DMZ, UPnP will automatically forward ports to it to the internet.
    + It is not recommended to set up UPnP and DMZs together.
    + This is because UPnP will forward ports to the devices in the DMZ, opening it to the internet.
    + It is more secure to manually setup port forwarding instead of using UPnP.

+ **Dynamic Host Configuration Protocol (DHCP)**
    + This is a protocol that automatically assigns IP addresses to devices.
    + Useful for home devices, but it will be problematic if used in a business environment.
    + This is because IP addresses are not static and will change.

+ **Custom Router Firmwares**: OpenWRT and DD-WRT
    + Both of these are open-source router firmwares based on Linux.
    + Busybox is a binary that is used in both of these firmwares.
    + Busybox is a compiled binary of many Linux commands (e.g. ```ls```, ```cat```, ```grep```, etc.).
    + If you want customization on your router, WRT firmwares on a NetGear router is recommended.

# Section 4: Firewalls

+ **Terms**:
    1. **Ingress/Inbound Filtering**: A firewall that filters traffic coming into the network.
    2. **Egress/Outbound Filtering**: A firewall that filters traffic going out of the network.
    3. **Access Control List (ACL)**: A list of rules that a firewall uses to filter traffic.
    4. **Implicit Deny**: A firewall rule that denies all traffic that does not match any rules.

+ **How NAT already protects us from ingress traffic**:
    + Internal devices do not get assigned external IPs, they only have internal IPs.
    + This is secure because an internal IP address is not routable on the internet.
    + Therefore, NAT already protects internal devices from ingress traffic.

+ **Examples of Egress filtering by firewall**:
    + Egress filtering means a firewall blocking traffic going out from the network to the internet.
    + An example of this is blocking malicious DNS requests from a device (that maybe part of a botnet).
    + Malicious egress traffic must be blocked because it bypasses NAT:
        1. If your device is part of a botnet, it will request the NAT to send traffic to a malicious server.
        2. If this happens, NAT will allow ingress traffic from the malicious server to the device.
        3. This way, the malicious server can send commands to the device over a remote shell bypassing the NAT.

+ **Network Segmentation**
    + This is basically dividing the network into separate sections so that lateral movement is impossible.

+ **Host/Network Based Firewalls**
    1. **Network-based firewalls**: These are firewalls that are placed between networks.

    2. **Host-based firewalls**: These are firewalls that are placed on devices.
        + These will block traffic based on the applications installed on a device on the network.
        + A whitelist of applications can be created only allowing them to access the network.
            + Eg: Only allow *Firefox* to access the network
            + Some examples of these firewalls are Windows Firewall & iptables
            + A disadvantage: malware could infect the whitelisted applications, and ride on them.

    3. **Virtual Firewalls**: These are firewalls that are placed on virtual machines.


# Section 5: Network Traffics, Architecture & Isolation

+ **Terms**:
    1. **Network Isolation**: This is the process of separating networks from each other.
    2. **ARP Spoofing**: This is the process of sending fake ARP packets to a device.

+ **ARP Spoofing**
    + **ARP** is a protocol that is used to map IP addresses to MAC addresses.
    + ARP is used by devices on a network to find the MAC address of another device.
    + It starts by a device sending an ARP request to the broadcast address asking for a MAC address associated with a specific IP address.
    + The device with the IP address will respond with its MAC address.
    + ARP spoofing occurs by sending fake ARP responses to a device.
    + ARP attack prevention such as `aprwatch` will prevent APR attacks.

+ Network Isolation
    + **Virtual Network Isolation**
        + This is the process of separating networks from each other using logical VLANs.
        + We can seperate networks logically by seperating into different VLANS:
            1. `192.168.1.0`: May be a very trusted network
            2. `192.168.1.1`: Might be an untrusted network
            3. `192.168.1.2`: Might be a guest network
        +  These are different networks logically seperated, and broadcast traffic will not be sent to other networks.
    + **Physical Network Isolation**
        + This is the process of separating networks from each other using physical switches.

+ **WiFi Weak Encryption & Their Fixes**
    1. **WEP**: This is a weak encryption protocol that can be cracked in minutes.
    2. **TKIP**: This is an encryption protocol used in WPA which is also weak.
    3. **CCMP (AES)**: This is used to replace TKIP in WPA2 which is strong.
    4. **Evil Twil Attack**: When a fake WiFi is created identical to the victim's WiFi.

# Section 6: Wireless & Wi-Fi Security

+ **WiFi best practices**
    1. Use less-common SSID names
        + If you use common SSID names such as `linksys` or `NETGEAR`, it can be vulnerable to rainbow table attacks.
        + This is because there are pre-computed rainbow tables for these SSIDs.

    2. Disable WPS
        + WPS is a feature that allows devices to connect to a network without a password.
        + This is vulnerable to brute-force attacks and can crack a WiFi in hours.

    3. Use RADIUS authentication in a business setting

    4. MAC filtering is not a security feature
        + MAC filtering is not a security feature because MAC addresses can be spoofed.

    5. If only needed, reduce the range of the WiFi's radio frequency to reduce the attack surface.

# Section 7: Network Monitoring for Threats

+ **Syslog** is a protocol that is used to send log messages to a server.
+ If a router has syslog enabled, it will send log messages to a syslog server IP that we specify.
+ **Wireshark, tcpdump, and tshark** are tools that can be used to capture network traffic.
    + Ideally, tcpdump will generate a netowrk dump file, and that is analyzed using Wireshark/tshark

# Section 8: How we are tracked online

+ Using pubic WiFis will put you in a big risk of being tracked online.
+ This is because the packets you send to the internet can bne sniffed by anyone on the network.
+ These are some ways how you can be tracked online:
    1. **IP Address**: This is a unique identifier assigned to your device.
    2. **3rd Party Connections**: This is when a website loads content from another website.
        + Eg: A website loading a YouTube video.
        + This is because the website will send your IP address to YouTube.
        + This means that YouTube will know your IP address.
    3. **HTTP Referrer**: When a webpage sends you to another webpage, it will send the URL of the previous webpage.
        + Eg: If you are on `example.com` and you click on a link to `example2.com`, `example2.com` will know that you came from `example.com`.
    4. **Cookies**: These can be used by websites to track you.
        + Cookies are small files that are stored on your device by websites.
        + These files can contain information such as your IP address, location, etc.
    5. **Super Cookies**: These are hard to remove cookies that can be used by ad companies to track you.
    6. **Geolocation**: This is the process of finding the location of a device.

+ **Profiling** can be done by ad companies to track you.
    + This is the process of collecting information about you and your device.
    + This can be done by using cookies, super cookies, and geolocation.
    + This information can be used to create a profile of you.
    + Then these profiles can be sold to other companies.

# Section 9: Search Engines & Privacy

+ **Google Analytics & Their Cookies**
    + Search engines like Google use Google Analytics, which is a tool that can be used to track you.
    + Google Analytics can be blocked using a browser extension called **uBlock Origin**.
    + **NID & SID** is a cookie that is used by Google to profile you for to send ads.

+ Alternatives to Google and non-privacy friendly search engines:
    1. **DuckDuckGo**
    2. **Startpage**
    3. **ixquick**
    4. **Disconnect**

# Section 10: Browser Security & Tracking Protection

+ It is recommended to use Firfox or other privacy-focused browsers instead of Google based browsers as they aren't privacy-focused and might try to profile you from their browsers itself.

+ Having the following functionalities in a browser is very risky and its better to disable these:
    1. Flash
    2. Java
    3. Silverlight

+ If possible, it is safe to disable Javascript as there were instances where dark web users were de-anonimized using vulnerable Javascript.
+ Do not install browser extensions that are not trusted.
+ There are options such as 'Browser in the box' which will run a browser in a virtual machine.
+ Its better to use Firefox profiles to isolate your user profiles.
+ Firefox best practices:
    1. Enable 'Do not track you'
    2. Enable 'Use Tracking Protection In Private Windows'
    3. Ensure that the Disconnect.me service is enabled on Firefox
    4. Enable 'Never remember history'
    5. Enable 'Block reported attack sites'
    6. Enable 'Block reported web forgeries'
    7. Block all HTTP websites

+ Recommended browser extensions:
    1. uBlock Origin
    2. uMatrix
    3. Disconnect, Ghostery, Request policy
    4. ABP (Ad Block Plus), Privacy badger, WOT
    5. No-script
    6. Policeman

+ **Browser fingerprinting**: A user can be uniquely identified by their browser fingerprint.
