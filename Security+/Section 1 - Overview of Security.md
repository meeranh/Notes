---
id: "Section 1 - Overview of Security"
aliases:
  - "CIA Triad"
tags: []
---
# CIA Triad
+ This is a model designed to **guide policies for information security** within an organization.
+ This includes protecting **three key concepts**:
  + **Confidentiality**: Making sure that information has not been disclosed to unauthorized people.
  + **Integrity**: Making sure that information can be changed only by authorized people.
  + **Availability**: Making sure that services are working when needed.

# AAA of Security
+ This is a framework that **makes sure that only users can access systems** while all their **operations on the systems they use are recorded**.
+ AAA refers to:
    1. **Authentication**: Verifying the identity of a user.
        + **Something you know** (password)
        + **Something you have** (token)
        + **Something you are** (biometrics)
        + **Something you do** (signature)
        + **Somewhere you are** (location)

    2. **Authorization**: Determines what a user is allowed to do.

    3. **Accounting**: Tracking what a user does with his/her access.
        + Log files are **used to track user activity**.
        + This is to **collect proof in case of a security incident**.

# Security Threats
+ The following are four of the main security threat categories an organization could face:
    1. **Malware**: Malicious software such as viruses/worms.
    2. **Unauthorized Access**: The bypassing of authorization and accessing sensitive data.
    3. **System Failure**: When a system crashes and is not available.
    4. **Social Engineering**: Tricking internal users into giving up sensitive information.

# Mitigating Threats
+ These are some of the ways to mitigate the threats mentioned above:
    1. **Physical Controls**: Real-world controls such as CCTV, locks, security guards, etc.
    2. **Technical Controls**: Controls implemented in software like encryption, firewalls, etc.
    3. **Administrative/Managerial Controls**: Organizational controls such as security policies, security awareness training, etc. This can be broken down into the two following types:
        a. **Procedural Controls**: Controls that the organization chooses to do on its own.
        b. **Legal & Regulatory Controls**: Controls required by laws and regulations (e.g. GDPR, HIPAA, etc.)

# Hackers
+ These are the following types of hackers that exist:
    1. **White Hat**: Break into systems for non-malicious purposes (e.g. penetration testing).
    2. **Black Hat**: Break into systems for malicious purposes.
    3. **Grey Hat**: Break into systems for both malicious and non-malicious purposes.
    4. **Blue Hat**: Hackers who aren't part of an organization but are invited to test their systems.
    5. **Elite**: Hackers who are highly skilled and find zero-day vulnerabilities.
    6. **Script Kiddie**: Hackers who are not skilled and use tools created by others to hack systems.

# Threat Actors
+ These are people who could cause harm to an organization:
    1. **Script Kiddie**: A hacker who is not skilled and uses tools created by others to hack systems.
    2. **Hacktivist**: A hacker who hacks for a cause, but illegal (anonymous).
    3. **Organized Crime**: A group of hackers who work together to hack systems
    4. **Advanced Persistent Threat (APT)**: A group of hackers who are highly skilled and have a lot of resources (often funded by governments).

# Threat Intelligent & Sources
## Threat Intelligence
+ Intelligence refers to **information that has been analyzed and is ready to be used**.
+ The following are some factors that make intelligence useful:
    1. **Timeliness**: The information is up-to-date *(information about a threat from 10 years ago is not useful)*.
    2. **Relevance**: The information is useful to our use case *(threat details about cars are not useful for a company that makes shoes)*.
    3. **Accuracy**: The information is correct *(information about a threat that is not true is not useful)*.
    4. **Confidence Level**: Is the information reliable? *(information from a trusted source is more reliable than information from an unknown source)*.
        + The MISP project can be used to evaluate a confidence of an intelligence source.
        + MISP gives a framework to assess confidence in both **information content** and **source reliability**.

## Sources
+ The following are three places you can get your information from:
    1. **Proprietary**: Payment is required for the information.
    2. **Closed Source**: Information is free but not available to the public.
    3. **Open Source**: Information is free and available to the public.
    4. **OSINT**: Gathering details about a target from open sources (Facebook, Twitter, etc.)

 ### Explicit Knowledge vs Implicit Knowledge
 1. **Explicit Knowledge**: Knowledge that is written down.
 2. **Implicit Knowledge**: Advice/instincs/experience from experts that is not written down.

# Threat Hunting
+ This is a method of which an organization can **proactively search for threats** without waiting for a security alert.
+ We will have to **assume that our intrusion detections systems failed**, and look through logs to find evidence of a breach.
+ The following is a guideline on how threat hunting can be done:
    1. **Establishing A Hypothesis**: Imagining a hypothetical attack that could occur.
    2. **Profiling Threat Actors & Activities**: What would be the attacker's motives, systems they would hack, etc.
    3. **Look through logs for evidence of the attack**
        1. Analyze network traffic logs.
        2. Analyze the executable process list.
        3. Analyze other infected hosts.
        4. Identify how the malicious process was executed.

## Benefits of Threat Hunting
1. **Improve detection capabilities**: You can refine your intrusion detection systems based on the threats you find.
2. **Integrate intelligence**: You can use previously gathered intelligence to help you find threats.
3. **Reduce attack surface**: You can identify where an attacker could attack and patch it.
4. **Block attack vectors**: You now identify the techniques an attacker could use and block them.
5. **Identify critical assets**: The most important assets can be identified and protected.

# Attack Frameworks
+ These are frameworks that can be used to perform an attack.

## 1. Lockheed Martin Kill Chain
+ The Lockheed Martin Kill Chain is a framework that lays out an attack plan. 
+ Security analysts can use this to hypothesize how an attacker would attack a target and protect the system accordingly.
+ The following stages are available in this framework:
    1. **Reconnaissance**: Getting to know the target and identifying vulnerabilities.
    2. **Weaponization**: Creating a payload to exploit the vulnerability.
    3. **Delivery**: Delivering the payload to the target.
    4. **Exploitation**: The target now executes the payload.
    5. **Installation**: The malicious payload is now installed on the target system.
    6. **Command & Control**: Remote access is granted to the attacker.
    7. **Action on Objectives**: Attacker now performs his initial attack objective.

## 2. MITRE ATT&CK Framework 
 + Unlike the Lockheed Martin Kill Chain, which was a linear framework, **this is a hollistic framework**.
 + It lays out all the **possible attacks/behaviours** a threat actor could have.
 + Organizations could use this to identify the techniques an attacker could use and protect against them.
 + A pre-ACTT&CK framework is also available, which is a framework that lays out the steps an attacker would take before the attack during reconnaissance.

## 3. Diamond Model of Intrusion Analysis
+ The Diamond Model is a structured method for analyzing and categorizing intrusions.
+ It focuses on the relationships between four core features of an intrusion: **Adversary**, **Capability**, **Infrastructure**, and **Victim**.
    1. **Adversary**: The person or group responsible for the intrusion.
    2. **Capability**: The tools, techniques, and procedures (TTPs) used by the adversary.
    3. **Infrastructure**: The infrastructure used by the adversary to launch the attack.
    4. **Victim**: The target of the attack.

+ These four features form the vertices of a diamond, hence the name.
+ It's a dynamic model that can be used in conjunction with other frameworks to provide a comprehensive view of the threat landscape.
