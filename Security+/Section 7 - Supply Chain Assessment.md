---
id: "Section 7 - Supply Chain Assessment"
aliases:
  - "Supply Chain Assessment"
tags: []
---

# Supply Chain Assessment
+ This is the practice of checking if the hardware/software that is being purchased from a vendor is secure.
+ **Due Diligence**: A legal identifying a subject has used best practices when setting, configuring, and maintaining a system. Some examples of due diligence are:
    1. **Properly resourced cybersecurity program**.
    2. **Security assurance & risk management processes are in place**.
    3. **Product support life cycle (will they provide security patches)**.
    4. **Security controls for confidential data is in place**.
    5. **Incident response/forensics capabilities**.
    6. **General & historical company information**.

+ **Trusted Foundry**: A way in which the Department of Defense (DoD) can ensure that the hardware they are purchasing is secure.

+ **Hardware Source Authenticity**: Ensuring that bought hardware is taper-proof and has not been tampered with (buying a router from Cisco is better than from Ebay)

# Root of Trust (ROT)
+ This is a **embedded cryptographic module** that is used to verify the integrity of a system (such as a TPM). These scan the boot files of a system and verifies their signatures.
+ **Anti-tamper** mechanisms such as FPGA (Field Programmable Gate Array) and PUF (Physically Unclonable Function) can be used to prevent tampering with the hardware.
    + These will wipe the data on the hardware if tampering is detected.

# Trusted Firmware
+ Firmware is software that is embedded into a hardware device such as the UEFI. There are a few ways in which insecure firmware can be prevented:
    1. **Secure Boot**: Checks if the programs booting up are signed by a trusted source.
    2. **Measured Boot**: Compares metrics such as time taken to boot to check if the system has been tampered with.
        + An attestation report is then made which is digitally signed by the TPM key.
    3. **eFUSE**: An electronic fuse storing cryptographic keys that will be blown if the system is tampered with.
    4. **Trusted Firmware Updates**: Verifying the checksum of the firmware before updating it.
    5. **Self-Encrypting Drive**: Firmware on the chip is encrypted.

# Secure Processing
+ It is possible that malicious bytecode is processed in the CPU, RAM, etc.
+ These are some of the ways in which secure processing can be ensured:
    1. **Processor Security Extensions**: Low-level CPU instructions (such as TXT or SGX) that enable secure processing.
    2. **Trusted Execution**: Invoke a TPM & secure boot attestation to run the trusted OS.
    3. **Secure Enclave**: Creates an encrypted container for sensitive data.
    4. **Atomic Execution**: Executions are either successfully performed, or not performed (rolled back if error).
    5. **Bus Encryption**: Data is encrypted before being sent over the bus (the main pathway for data transfer in the computer system).
