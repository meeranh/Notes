---
id: "Section 8 - Virtualization"
aliases:
  - "Virtualization"
tags: []
---

# Virtualization
+ The process of running multiple operating systems on a single physical machine
+ There are two types of virtual machines:
    1. **System Virtual Machine**: A virtual machine that runs a complete operating system.
    2. **Processor Virtual Machine**: A virtual machine that runs a single application.

# Hypervisor
+ This manages the system resources that is sent to the virtual machines to use.
1. **Type 1 Hypervisor**: A hypervisor that runs directly on the host's hardware (Hyper-V).
2. **Type 2 Hypervisor**: A hypervisor that runs as a software on the host's operating system (VirtualBox).
3. **Application Containerization**: A single shared kernel that runs multiple applications with different user-data.

# Threats to VM
1. **VM Escape**: When an attacker accesses the host machine from the virtual machine by bypassing the hypervisor.
2. **Data Remnants**: When de-provisioning a virtual machine, the data is not completely removed from the host machine.
3. **Privilege Escalation**: When a user grants themselves more privileges than they should have and gains access to the host machine.
4. **MITM during live migration**: When an attacker intercepts the data being transferred from one host machine to another.

# Securing VMs
1. **Keep the hypervisor up-to-date**
2. **Limit connectivity between the virtual & host machine**
3. **Ensure proper patch management to avoid virtualization sprawl**
    + **Virtualization Sprawl**: When there are too many virtual machines without proper management.
4. **Enable disk-encryption of virtual disk images**
