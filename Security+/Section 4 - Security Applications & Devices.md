---
id: "Section 4 - Security Applications & Devices"
aliases:
  - "Software Firewalls"
tags: []
---

# Software Firewalls
1. **Personal Firewalls/Host-based Firewalls**: Protects a single computer from unwanted traffic.
    + Windows has Windows Firewall *(accessed by `wf.msc`)*
    + Mac has PF & IPFW
    + Linux has iptables

# Intrusion Detection System (IDS)
+ Device/software that analyzes the data packets going through a network to identify an incident.
+ They only can give out alerts of potential attacks, they cannot stop them.
+ Ideally, all system logs should be sent to a syslog server.
+ There are two types of IDSs:
    1. **Host-based Intrusion Detection Systems (HIDS)**: Monitors a single machine for suspicious activity.
    2. **Network-based Intrusion Detection Systems (NIDS)**: Monitors an entire network for suspicious activity.

+ The following is how IDSs detect intrusions:
    1. **Signature-based**: Checks the data packet against a database of known attack signatures.
    2. **Policy-based**: Checks packets against company policy *(eg: Cannot use Telnet)*.
    3. **Anomaly-based**: If incoming traffic is different from the normal traffic (an estabished baseline), it is flagged as suspicious.

+ The accuracy of an IDS's alerts must be balanced with the number of false positives and false negatives. The following are the types of alerts an IDS can generate:
    1. **True Positive**: Malicious activity is identified as an attack.
    2. **True Negative**: Legitimate activity is identified as legitimate traffic.
    3. **False Positive**: Legitimate activity is identified as an attack.
    4. **False Negative**: Malicious activity is identified as legitimate traffic.

# Pop-up Blockers
+ Pop-ups can be blocked using the following software:
    1. **Ad-blocker**: Blocks ads on websites.
    2. **Content Filters**: Blocks certain sites or content such as Javascript, images, etc.

# Data Loss Prevention (DLP)
+ Monitors data at rest, in use, or in transit for unauthorized access or exfiltration.
+ There are four main types of DLPs:
    1. **Endpoint DLP System**: Monitors read/write access to data on a single machine.
    2. **Network DLP System**: Monitors data in transit on a network.
    3. **Storage DLP System**: Monitors data at rest in server rooms.
    4. **Cloud DLP System**: Protects data stored in the cloud.

# Security The BIOS/UEFI
+ BIOS is the first thing that runs when a computer is turned on.
+ The BIOS can be protected by the following:
    1. **Flash the BIOS**: Keeping the BIOS up-to-date by flashing new BIOS updates.
    2. **Use a BIOS password**: Use a strong password to change BIOS settings.
    3. **Configure the BIOS boot order**: Make sure that the right devices are being booted on startup.
    4. **Disable external ports & devices**: Unused ports and devices should be disabled.
    5. **Enable secure boot**: Will load the public key in the TPM to verify the OS.

# Security Storage Devices
+ The follwoing are some ways in which you can secure your storage devices:
    1. **Encrypt data on your storage device**: Data that is stored on devices like thumb drives should be encrypted.
    2. **Removable media controls**: Technical/administrative controls such as disabling USB ports.
    3. **Use a Network Attached Storage (NAS)**: A NAS is a storage device that is connected to a network and can be accessed by multiple users. It uses RAID which gives high availability.
    4. **Use a Storage Area Network (SAN)**: A collection of NASes.

+ The following can be done to protect NAS and SAN:
    1. **Use data encryption**
    2. **Use proper authentication**
    3. **Log NAS accesses**

# Disk Encryption
+ There are two kinds of disk encryptions:
    1. Software based disk encryption
        + **Self-Encrypting Drive (SED)**: A storage device that performs whole disk encryption using embedded hardware.
            + In Windows, **BitLocker** can be used to achieve this. It uses the TPM (Trusted Platform Module) on the motherboard to get an encryption key.
            + These uses **AES** for encryption.
    
    2. Hardware based disk encryption
        + **Hardware Security Module**: Physical devices that generate, store, and manage encryption keys.

# Endpoint Analysis
+ The following tools can be used for endpoint analysis:
    1. **Anti-Virus**: Software that detects and removes malware.
    2. **Host-based IDS/IPS**: Monitors/prevents a single machine for suspicious activity.
    3. **Endpoint Protection Platform (EPP)**: A software that provides anti-virus, anti-malware, and host-based IDS/IPS, firewall, DLP, etc.
    4. **Endpoint Detection & Response (EDR)**: Collects system logs and provides early indications of an attack.
    5. **User & Entitiy Behaviour Analytics (UEBA)**: Provides a baseline of normal user behaviour and alerts when there is a deviation from the baseline in user accounts/computer hosts. An example is Splunk.
