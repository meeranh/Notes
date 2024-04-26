# Virtual Machines In Azure
+ Much like normal virtual machines, Azure virtual machines are essentially a computer within a computer.

# Creating A Basic VM In Azure
+ Go to 'Virtual Machines' in [Azure Portal](https://portal.azure.com)
+ Click 'Add'
+ Select your [[Resource Groups|Resource Group]].
+ Fill in your machine name.
+ Select your [[Regions & Zones]] from the 'Availability Options' drop-down.
+ You can also choose an [[Availability Set]].
+ Choose your operating system from the "Image" drop-down.
+ Choose the CPU and memory size from the "Size" drop-down.
+ For authentication, you can use a password or SSH public key.
    + If you were to choose an SSH public key, you would need to generate one.
+ Choose the ports you want to open.
+ Click 'Review + Create'.

### Adding Disks To Your VM
+ You can go to the 'Disks' tab and add a new disk to your VM.
+ From the 'OS Disk Size' drop-down, you can choose the size of the disk.
+ Then, in the 'OS Disk Type' drop-down, you can choose:
    1. **Standard HDD**: These are the cheapest but slowest (95% availability)
    2. **Standard SSD**: These are faster than HDD but slower than Premium SSD (99.5% availability)
    3. **Premium SSD**: These are the fastest but the most expensive (99.9% availability)

# Cloud Init Scripts
+ Lets say that for every VM you deploy, you want to update, upgrade and install a package.
+ The way to do this is to use a cloud-init script.
1. When creating a VM, go to the "Advanced" tab.
2. In the "Custom Data" field, you can paste your cloud-init script.
+ This is basically a bash script (if Linux) that will run when the VM is created.
```bash
#!/bin/bash
apt-get update
apt-get upgrade -y
apt-get install -y apache2
```

+ Paste this into the "Custom Data" field and click 'Review + Create'.

# Connecting To An Linux VM via SSH
+ The SSH private key you got from Azure does not have the right permissions.
+ To change the permissions to an usable state, run the following command:
```bash
chmod 400 your-downloaded-private-key.pem
```
+ Then, you will be able to connect to your VM via SSH:
```bash
ssh -i your-downloaded-private-key.pem your-username@your-vm-ip
```

# Connecting via Azure Cloud Shell
+ Azure Cloud Shell is basically just a terminal in your browser.
+ If you don't want to use your own terminal, you can use Azure Cloud Shell.
+ You will have to upload your SSH private key to Azure Cloud Shell and run the above commands inside it.

# Azure VM Key Concepts
### Image
+ These are the operating systems that you can run on your newly created virtual machine.

### VM Sizes
+ These are combinations of CPU and memory that you can choose from.
+ Think of these like packages of CPU/Memory/GPU/Storage.
+ You can choose the one that fits your needs.

### VM Family
+ Many VM sizes are grouped into families.
+ The VM sizes with very low CPU and memory are in the "A" family.
+ The VM sizes with high CPU and memory are in the "D" family.
+ The VM sizes with GPU are in the "N" family.
+ As such, you can choose the family that fits your needs.
+ It is easier to first choose a family and then choose a size from that family.
