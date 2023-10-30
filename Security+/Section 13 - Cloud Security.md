---
id: "Section 13 - Cloud Security"
aliases:
  - "Cloud Security"
tags: []
---

# Cloud Security
+ Using virtualization on remote servers to provide services.
**Virtual Desktop Infrastructure (VDI)**: A user is given a virtual desktop that is hosted on a remote server that is destroyed after some time (1 day).
    + **Secure Enclaves**: A secure area of memory that is encrypted and only accessible by the CPU.
    + **Secure Volumes**: A volume is decrypted when being used, but encrypted when it is not in use.

# Cloud Types
1. **Public**: Any user on the internet can use this cloud (ex: Google Drive).
2. **Private**: A company creates its own personal cloud environment.
3. **Hybrid**: A combination of public and private cloud.
4. **Community**: Two or more organizations share a cloud environment.

# As A Service
1. **SaaS (Software as a Service)**: Provides software applications over the internet on a subscription basis (ex: Gmail, Microsoft 365).

2. **IaaS (Infrastructure as a Service)**: Provides virtualized computing resources over the internet. Users can rent IT infrastructure such as servers and storage (ex: Amazon Web Services).

3. **PaaS (Platform as a Service)**: Provides a platform and environment for developers to build applications and services (ex: Google App Engine).

4. **SECaaS (Security as a Service)**: Delivers security services and solutions over the cloud on a subscription basis (ex: McAfee Cloud Security).

# Cloud Security
+ These are some good practices a cloud practictioner can follow to be safe on cloud environments:
1. Impost proper access controls to data as data is collated (stored on the same server).
2. Have the latest updates to the VM's software.
3. Have redundancy of data.
4. Use proper security policies.
5. Data should always be encrypted to prevent data remenents after deprovisioning.

# Defending Servers
1. **File Servers**: Used to store files and folders.
    + Use data encrytion on these servers.
    + Keep the firmware and OS up to date.

2. **Email Servers**: Routes emails to the correct destination.
    + Keep the firmware and OS up to date.
    + Use spam filters to prevent spam emails.
    + Use anti-malware software to prevent malware from being sent through emails.

3. **Web Servers**: The server which hosts the public webpages.
    + Keep the firmware and OS up to date.
    + Put the web server in a DMZ.

3. **FTP Servers**: Servers that store data and allow users to download and upload files.
    + Keep the firmware and OS up to date.
    + Use TLS to encrypt the data.

4. **Domain Controller (Active Directory)**: A server that manages the users and computers in a domain.
    + Keep the firmware and OS up to date.
    + Use a strong password policy.

# Cloud-Based Infrastructure
 + **Virtual Private Cloud (VPC)**: When resources are deployed on a private cloud and given to public consumers.
    + Typically used to host web servers.

+ When moving to the cloud, there could be problems such as:
    1) **Vendor Lock In**: When a very big portion your data is now stored on the cloud that you have no choice but to use this vendor.
    2) **Compliance**: When a company has to follow certain rules and regulations when storing data.

# Cloud Access Security Broker (CASB)
+ This is a service that sits between the cloud provider and the cloud consumer.
+ I shows what the users are doing on the cloud and can prevent certain actions from happening.
    + **Forward Proxy**: A proxy that sits in front of the cloud provider.
    + **Reverese Proxy**: You cant connect to the cloud network without the proxy
    + **Application Programmable Interface (API)**: Using HTTP requests to access the cloud.

# Function as a Service (FaaS)
+ The ability to run services on the cloud without a VM.
+ Another name for this is **Serverless**.
+ Some examples are Lambda, Azure Functions, and Google Cloud Functions.

# Cloud Threats
1. **API**
    + An API should be protected by HTTPS.
    + Input validation has to be done on the server-side.
    + Errors should be handled properly (detailed error messages should not be shown to the user).
    + Implement throttling/rate-limiting on the API

2. **Improper key management**
    + Protect your API keys so that unauthorized access is not possible.
    + Do not hardcode your keys onto the source code of your applications.
    + Delete unnecessary keys when not needed

3. **Insufficient Logging**
    + Logs on the cloud should be copied to a different location.
        *eg: Cloudwatch & DataDog for AWS Lambda*

4. **Default Container Settings**
    + Leaving the default settings on containers like blobs or buckets might result in data leakage.

5. **No origin policies**
    + Implementing a CORS (Cross Site Origin Resource Sharing) policy can prevent unauthorized access to your cloud resources.
