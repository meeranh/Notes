---
id: "Section 5 - Mobile Device Security"
aliases:
  - "Securing Wireless Devices"
tags: []
---

# Securing Wireless Devices
+ Most wireless devices are either one of the following:
    1. **WiFi**: *Always use WPA2 (WiFi Protected Access 2) which relies on AES*.
    2. **Bluetooth**: *When paired, a shared link key is exchanged to encrypt data*.
+ If confidentiality is a requirement, try to **stay away from wireless devices**.
+ However if wireless devices are required, **always check if encryption is done using AES and a strong key**.

# Mobile Malware
+ Here are some steps to prevent mobile malware:
    1. **Ensure your mobile is patched and updated**
    2. **Install apps from the App Store/Play Store**
    3. **Do not jailbreak/root devices**
    4. **Do not use custom firmware/custom ROM**

# SIM Cloning & ID Theft
+ **SIM**: This is the integrated circuit that stores the IMSI (International Mobile Subscriber Identity) and its relavant key.
+ **SIM cloning** is the process of copying the IMSI and its key to another SIM card. This was possible in v1 but harder in v2.
+ **ID Theft** can occur if you make your email/phone number public, which could be taken over by an attacker.
+ The mitigation for these is to **be careful where you post your phone number**.

# Bluetooth
+ **Bluejacking**: Sending unsolicited messages to other Bluetooth devices.
+ **Bluesnarfing**: Listening to unauthorized information from wireless Bluetooth devices.
+ The following are some methods in which these attacks can be prevented:
    1. Change the default pairing key (PIN)
    2. Turn off Bluetooth when not in use.
    3. Disable discoverable mode if you are using Bluetooth.

# Mobile Device Theft
+ Devices can be stolen, these are some ways in which data in these devices can be protected in case of theft:
    1. *Always keep a backup of device data.*
    2. *Use full-disk encryption on the device.*
    3. *Set up location tracking on the device.*
    4. *Use the 'Remote Lock' feature to lock the phone until found.*
    5. *Use the 'Remote Wipe' feature to delete all data on the phone remotely.*

# Security of Apps
+ The following precautions should be taken to make sure that you are safe from malicious applications:
    1. Install applications only from the official Play Store/App Store.
    2. Use a trusted browser such as Google Chrome
    3. Always visit HTTPS websites
    4. Turn location off for apps that do not need your location.
    5. Disable geotagging (metadata on images).

## Enterprise Mobility
+ This is the fact that employees will be using their devices or company devices to access company resources.
+ **Mobile Device Management (MDM)** can be installed on devices to manage company devices remotely to prevent the installation of apps without the company's permissions.

# BYOD (Bring Your Own Device)
+ To combat the security issues of BYOD, the following can be done:
    1. **Storage Segmentation**: This is the process of separating personal and company data on the device.
    2. **Mobile Device Management (MDM)**: This is the process of managing company devices remotely.
        + **CYOD (Choose Your Own Device)** is the process of allowing employees to choose their own devices from a list of approved devices.
        + Usually, MDM is enabled on CYOD devices. This is because employees might be reluctant to enable MDM on BYOD devices.

# Hardening Mobile Devices
+ The following are some ways to improve the security of your mobile devices:
    1. Update your device to the latest software version.
    2. Install anti-virus.
    3. Train users on proper security on the use of the device.
    4. Only install applications from the App Store/Play Store.
    5. Do not root/jailbreak your device.
    6. Only use v2 SIM cards
    7. Disable all features that you do not use *(WiFi, Bluetooth, NFC, etc)*.
    8. Turn on encryption for voice & data.
    9. Use strong passwords or biometrics for authentication.
    10. Don't allow BYOD.
