---
id: "Section 12 - Perimeter Security"
aliases:
  - "Firewalls"
tags: []
---

# Firewalls
+ These are the functions of firewalls:
1. **Packet Filtering**: Monitoring/blocking the packets passing through the firewall based on allow/reject rules
    1. *Stateful Packet Filtering*: Does not keep track of the details of previous incoming packets.
    2. *Stateful Packet Filtering*: Keeps track of details of outgoing packets and can make decisions based on the context of the packet.

2. **NAT Filtering**: Filter based on ports/connection type such as TCP & UDP
3. **Application Layer Gateway (ALG)**: Filters based on application layer protocols such as HTTP, FTP, SMTP, etc.
4. **Circuit Level Gateway**: Filters traffic based on the session layer after the TCP handshake is complete.
5. **MAC Filtering**: Access contol lists can be written to allow or deny traffic based on the MAC address of the source or destination.
    + *Explicit Allow*: A defined rule for an IP address range created to allow traffic.
    + *Explicit Deny*: A defined rule for an IP address range created to deny traffic.
    + *Implicit Deny*: If no rules are defined for an IP address, they are denied by default.

+ **Web Application Firewalls (WAFs)**: These are firewalls designed to detect and block traffic that attempt XSS, SQL Injection, and other kinds of web-based attacks.

# Proxy Servers
+ These are servers that sit in between you and your internet traffic sources (like an intentional man-in-the-middle)
+ There are 4 types of proxy servers:
1. **IP Proxy**: Hides the IP address of the client when browsing
2. **Caching Proxy**: Stores a copy of the web page on the proxy server to speed up future requests
3. **Internet Content Filter**: Blocks access to certain websites
4. **Web Security Gateway**: A combination of the above 3 and also scans for viruses and other features.

# Honeypots & Honeynets
+ **Honeypot**: A computer/file that is intentionally left vulnerable to attract attackers
+ **Honeynet**: A network of honeypots

# Data Loss Prevention (DLP)
+ A system that monitors data in use, data in motion, and data at rest to prevent data loss.
+ It can read data in emails, files, and other data sources to detect sensitive data and prevent it from being sent out of the network.

# NIDS vs NIPS
+ **Network Intrusion Detection System (NIDS)**: A system that monitors network traffic and *only alerts* when suspicious activity is detected.

+ **Network Intrusion Prevention System (NIPS)**: A system that *monitors and blocks* suspicious traffic.
    + This can also be used as a *protocol analyzer*, which is a tool that captures and analyzes packets.

# Unified Threat Management (UTM)
+ Also known as **Next Generation Firewalls (NGFW)**
+ This is a combination of multiple security controls into a single device.
+ Services such as NIDS, NIPS, DLP, and others can be combined into a single device.
