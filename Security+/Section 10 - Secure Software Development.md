---
id: "Section 10 - Secure Software Development"
aliases:
  - "Software Development"
tags: []
---

# Software Development
+ These are the phases of **Software Development Life Cycle (SDLC)**
1. **Planning & Analysis**: Gathering information & requirements from the client.
2. **Software/System Design**: Designing the software/system with diagrams.
3. **Implementation**: Coding the software/system.
4. **Testing**: Checking if the software/system works as intended.
5. **Integration**: Connecting the software/system to other systems.
6. **Deployment**: Hosting the software/system on a server.
7. **Maintenance**: Updating the software/system to fix bugs & add features.

+ **Waterfall Model**: When development teams finish SDLC phases one-by-one in a linear fashion.
+ **DevOps**: Development teams and operations teams work together.

# SDLC Principles
+ **Threat Modeling**: Security analysts potential vectors of attack.
+ These are some best practices to build secure software:
    1. **Least Privilege**: Giving users the least amount of privileges to do their job.
    2. **Defense in Depth**: Layer security controls to protect against multiple threats.
    3. **Never Trust User Input**: Always sanitize user input to prevent code injection.
    4. **Minimize Attack Surface**: Reduce the use of unneeded plugins/add-ons etc.
    5. **Create Secure Defaults**: Set secure settings by default.
    6. **Authenticity & Integrity**: Applications should be code-signed.
    7. **Fail Securely**: Have proper error handlings sytems in place.
    8. **Fix Security Issues**: Patch a vulnerability immediately after discovery.
    9. **Rely on Trusted SDKs**: SDKs are libraries that developers use to build software.

# Testing Methods
1. **Black-Box Testing**: Testing software without knowing how it works.
2. **White-Box Testing**: Tester has code access to the software.
3. **Gray-Box Testing**: Tester has some code access to the software.
+ These are best practices to use when programming: 
    1. **Use Structured Exception Handling (SEH)**: Provides controls on what to do if there is an exception/error.
    2. **Use Input validation**: The process of checking if user input is valid.
+ These are code analysis methods: 
    1. **Static Analysis**: Analyzing code without running it.
    2. **Dynamic Analysis**: Observing the behaviour of the software during runtime.

# Software Vulnerabilities
1. **Backdoors**: A secret way to bypass authentication.
2. **Directory Traversal**: A way to access files outside of the web root directory.
3. **Arbitary Code Execution**: A way to execute code on a system.
4. **Remote Code Execution**: A vulnerability to execute code on a remote system.
5. **Zero Days**: A vulnerability that is unknown to the public.

# Buffer Overflows
+ When a user enters data that is greater that the allocated buffer.
+ **Stack**: A memory location that stores data.
+ **Smash The Stack**: When a hacker overflows the stack with data.
+ **Address Space Layout Randomization** is a mitigation technique that randomizes the location of the stack.

# XSS & XSRF/CSRF
+ These are the following **types of XSS**:
    1. **Stored/Persistent**: When malicious code is stored on the server.
    2. **Reflected**: When malicious code is reflected back to the user (using URL parameters).
    3. **DOM-Based**: Attempts to exploit the browser.

+ **XSRF/CSRF** is sending malicious requests to a website on behalf of the user's session.
+ The following mitigations can be taken for these attacks:
    1. **XSS**: Perform output encoding.
    2. **XSRF/CSRF**: Use a CSRF token.

# SQL Injection
+ This is the processing of malicious user-provided SQL data that gets processed.
+ To prevent this, input validation and least privilege should be used.

# XML Vulnerabilities
+ These are some of the XML vulnerabilities:
    1. **XML Bomb (Billion Laugh Attack)**: Encodes XML entities that expand to exponential sizes, causing resource exhaustion.
    2. **XML External Entity (XXE)**: Requests a local resource (Like SSRF).

# Race Conditions
+ When two processes try to access the same resource at the same time due to dereferencing.
+ **Time to Check to Time to Use**: This is when the a resource is checked and read at different times (ideally, they both should be at the same time).
+ We can prevent this by the following methods:
    1. Develop applications to not process things sequentially if possible.
    2. Implement a locking mechanism to provide app with exclusive access.

# Design Vulnerabilities
1. **Insecure Components**: Pulling code from external sources (Stackoverflow, SDKs, etc).
2. **Insufficient Logging & Monitoring**: Does not properly log errors/exceptions.
3. **Weak or Default Configurations**: When defauls such as *admin:admin* are set.
+ To prevent these, use scripted installations & baseline configuration templates to secure applications during installations.
