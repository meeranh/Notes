---
id: Section 15 - Network Attacks
aliases:
  - Network Attacks
tags: []
---

# Network Attacks
+ Some of the most common types of network attacks are:
    1. Denial of Service
    2. Spoofing
    3. Hijacking
    4. Replay
    5. Transitive Attacks
    6. DNS Attacks
    7. ARP Poisoning

# Ports & Protocols
+ Inbound Port: Open to other services waiting for a connection
+ Outbound Port: When a service initiates a connection to another service

### Port Categories
1. **Well-Known Ports**: Port 0 - 1023 (Assigned by IANA)
2. **Registered Ports**: Port 1024 - 49,151 (Registered by vendors for specific applications)
3. **Dynamic/Private Ports**: Port 49,152 - 65,535 (Can be used by any application)

# Memorization of Ports

| Service | Port |
|---------|------|
| FTP | 21 |
| SSH, SCP, SFTP | 22 |
| Telnet | 23 |
| SMTP | 25 |
| DNS | 53 |
| TFTP | 69 |
| HTTP | 80 |
| Kerberos | 88 |
| POP3 | 110 |
| NNTP | 119 |
| RPC | 135 |
| NetBioS | 137-139 |
| IMAP4 | 143 |
| SNMP | 161 |
| SNMP Trap | 162 |
| LDAP | 389 |
| HTTPS | 443 |
| SMB | 445 |
| SMTP with SSL | 465/587 |
| Syslog | 514 |
| LDAP SSL/TLS | 636 |
| iSCSI | 860 |
| FTPS | 989/990 |
| IMAP SSL/TLS | 993 |
| POP3 SSL/TLS | 995 |
| MSSQL | 1433 |
| RADIUS (Alternate) | 1645/1646 |
| L2TP | 1701 |
| PPTP | 1723 |
| RADIUS | 1812/1813 |
| FCIP | 1937 |
| iSCSI Target | 3260 |
| RDP | 3389 |
| Diameter | 3868 |
| Syslog SSL/TLS | 6514 |

# Denial of Service
+ Any attack that aims to make a server unavailable.
+ There are 5 types of Denial of Service attacks:
    1. **Flood Attacks**: Sending a lot of packets to a server to overwhelm it.
        + **Ping Flood**: Sending a lot of ping requests to a server to overwhelm it.
        + **Smurf Attack**: Sending a lot of ping requests to a broadcast address to overwhelm a server.
        + **Fraggle Attack**: Uses UDP instead of ICMP to overwhelm a server.
        + **SYN Flood**: Sending a SYN but not completing the handshake. This makes the server wait for a response that never comes.
        + **XMAS Attack**: Set the FIN, PSH, and URG flags to crash the server.

    2. **Ping of Death**: Send an oversized ping packet to a server to crash it.
    3. **Teardrop Attack**: Send a packet with overlapping fragments to a server to crash it.
    4. **Permenant DoS**: Attempts to overwrite the firmware of a device to make it unusable.
    5. **Fork Bomb**: Creates a large number of processes to overwhelm a server.

# DDoS
+ **Distributed Denial of Service (DDoS)**: When multiple systems are used to perform a DoS attack.
    + **DNS Amplification Attack**: Uses a botnet to send a lot of DNS responses to a DNS query.

+ **Blackholing/Sinkholing**: Identifying attacker IPs and routing them to a null interface to stop DDoS attacks.

# Spoofing
+ When an IP address is forged to impersonate another system.
    + **IP Spoofing**: When an attacker uses a different IP address to impersonate another system.
    + **MAC Spoofing**: When an attacker uses a different MAC address to impersonate another system.

+ **Mitigation**: Using proper authentication and encryption will prevent spoofing from happening

# Hijacking
+ There are 8 kinds of hijacking attacks:
    1. **Session Theft**: Attacker guesses a session ID to impersonate a user.
    2. **TCP/IP Hijacking**: Attacker takes over a TCP session.
    3. **Blind Hijacking**: Injects data into a communication stream and checks if the injection was successful.
    4. **Clickjacking**: Attacker tricks a user into clicking a link that performs an action.
    5. **Main-in-the-Middle**: Attacker intercepts a communication between two parties.
    6. **Main-in-the-Browser**: Modifies web pages in the browser.
    7. **Watering Hole**: Malware placed in a very commonly visited website.
    8. **Cross-site Scripting**: Injects malicious code into a website.

# Replay Attack
+ When an attacker captures a packet, modifies it, and sends it back to the server.

# Transitive Attacks
+ **Transitive Logic**: A = B = C
+ A transitive attack is where a machine A trusts machine B and machine B trusts machine C.
+ Therefore, I can attack machine C and machine C will automatically be trusted by A

# DNS Attacks
+ There are 4 kinds of know DNS attacks
    1. **DNS Poisoning**: When an attacker changes the IP address of a domain name in the machine's cache.
        + **Mitigation**: DNSSEC is a protocol that uses digital signatures to prevent DNS poisoning.
    2. **Unauthorized Zone Transfers**: When an attacker requests replication of the DNS information on a system to attack later.
    3. **Altered Host Files**: Modifying the server's host file to reroute traffic to a malicious website.
        + **Pharming**: Reroutes traffic to a malicious website.
    
    4. **Domain Name Kiting**: When an attacker registers a domain name for a short period of time and then cancels it to avoid paying for it.

# ARP Poisoning
+ ARP translates IP addresses to MAC addresses before routing
+ Poisoning the ARP will cause the server to route traffic to the wrong MAC address. 
