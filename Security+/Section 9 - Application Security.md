---
id: "Section 9 - Application Security"
aliases:
  - "Web Browser Security"
tags: []
---

# Web Browser Security
+ Let others install patches, then wait a few days to see if there are any issues.
+ The following are some of the **general security recommendations for web browsers**:
    1. **Implement Policies**: Block access to certain websites.
    2. **Train your users**: Educate users to use HTTPS, don't clock on ads, etc.
    3. **Use proxy/content filter**: Proxies cache the data from websites and serve it to the users. This can help with performance and security.
    4. **Prevent malicious code**: Block the execution of Javascript, Flash, etc.

# Web Browser Concerns
1. **Cookies**: User authentication data that is stored on the user's computer, which can be leaked out.
    1. **Tracking Cookies**: Cookies that track the user's activity on the web.
    2. **Session Cookies**: Cookies that are used to keep the user logged in.
2. **Locally Shared Object**: Cookies for Adobe Flash
3. **Add-Ons**: Extensions that can be installed on the web browser to add more functionality.
4. **Advanced Security Options**: Browser configurations such as SSL/TLS settings, local storage, etc.

# Securing Applications
+ These are some best practices to prevent attacks from productivity applications such as Microsoft Word.
    1. **Use passwords to protect your files**.
    2. **Disable macros**.
    3. **Use digital certificates to verify the authenticity of the file**.
    4. **Set up User Access Control (UAC) to prevent unauthorized access**.
