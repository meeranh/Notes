---
id: "Section 6 - Hardening"
aliases:
  - "Hardening"
tags: []
---

# Hardening
+ The process of **adding layers of security** to reduce the attack surface.

# Unnecessary Applications
+ **Least Functionality**: Configuring a workstation/server to provide only essential applications/services.
+ A secure **Baseline Image** is recommended when installing an OS on a new workstation/server.
+ The **SCCM** (System Center Configuration Manager) is a Windows tool that can manage installations of applications over an enterprise.

# Restricting Applications
+ **Application Allowlisting**: Only applications on the allowlist are allowed to run.
+ **Application Blocklisting**: Only applications on the blocklist are not allowed to run.

# Trusted Operating System (TOS)
+ An operating system that meets the requirements set forth by government and has multilevel security.
+ **Windows 7 & newer** are TOS compliant.
+ **Mac OS X 10.6 & newer** are TOS compliant.
+ **Red Hat Enterprise Server** is TOS compliant.
+ **FreeBSD with TrustedBSD** is TOS compliant.

# Updates & Patches
+ **Hotfix/Patch**: A single update to a software application.
+ **Security Update**: A hotfix/patch that addresses a security vulnerability.
+ **Critical Update**: A critical issue (not a security issue) that needs to be addressed.
+ **Service Pack**: A collection of hotfixes/patches.
+ **Windows Update**: A service that provides updates to Windows OS.
+ **Driver Update**: An update to the driver software that allows a hardware device to communicate with an OS.

# Patch Management
+ **Patch Management** is the process of planning, testing, implementing, and auditing software patches.
    + **Planning**: Looking for patches that need to be applied and checking if they are compatible.
    + **Testing**: Running the patches in a test environment to see if they cause any issues.
    + **Implementing**: Applying the patches to the production environment.
    + **Auditing**: Making sure that the patches have been installed properly.

# Group Policy
+ A set of rules/policies that can be applied to a set of users/computer accounts within the operating system.
+ Usually, a **security template** is defined and loaded into the **Group Policy Object (GPO)**.
    + **Security Template**: A file that contains a set of security configurations.
    + **Group Policy Object**: A collection of Group Policy settings.

+ **Baselining** is first done to find what 'normal' usage is, and then the GPO is applied to the system to prevent any 'abnormal' activity.

# File Systems & Hard Drives
1. **FAT32**: An old Windows filesystem which is less secure.
2. **NTFS**: The default Windows filesystem which is more secure.
3. **ext4**: The default Linux filesystem.
4. **APFS**: The default Mac OS filesystem.

+ These are some ways to recover from hard-drive failure:
    1. Remove temporary files using Disk Cleanup
    2. Periodic filesystem checks
        + **Windows**: `chkdsk`, `sfc`
        + **Linux**: `fsck`
        + **Mac OS**: First Aid
    3. Defragment your disk drive
    4. Back up your data
    5. Use & practice restoration techniques
