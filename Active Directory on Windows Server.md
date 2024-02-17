# Section 2: Active Directory Overview & Terminology

+ **Term**
    1. **Domain**: Collection of objects (User accounts, Computers, Groups, etc).
    2. **Domain Controller**: The server with Active Directory installed.
    3. **Tree**: A collection of domains.
        + We might have a main domain named *example.com*.
        + We also might have a child domain named *child.example.com*.
        + A second child domain named *child2.example.com*.
        + This is basically a tree where example.com is the parent domain.
    4. **Forest**: A collection of trees.
        + We might have a domain named *example.com*.
        + We also might have a domain named *example2.com*.
        + These two might have a trust relationship between them.
        + Therefore, these two domain trees can be connected and called a forest.
        + This way, one domain can access a shared resource of another domain.
    5. **Global Catalog**: A list of all objects in the forest.
        + Imagine we go to Outlook and want to get details about a person in another tree.
        + The domain controller will query the entire tree and make a list of all the users/objects.
        + This list is known as the *Global Catalog*.
    6. **Read Only Domain Controller**: A domain controller that cannot be modified.
        + This is useful for branch offices or as a backup domain controller
    7. **DNS**: Resolve a domain name to an IP address.

# Section 3: Active Directory Installation
+ Steps to set up an Active Directory Domain Controller:
    1. Go to Server Manager.
    2. Set a proper computer name by:
        1. Select 'Computer Name'.
        2. Select 'Change'.
        3. Rename the server's name.
    3. Select 'Manage' and choose 'Add Roles and Features'.
    4. Select 'Next' on the 'Before you continue' window.
    5. Select 'Role-based of feature-based installation'.
    6. Select your machine & select 'Next' if your static IP is correct, else refresh the list.
    7. Select 'Active Directory Domain Services' & its additional features to setup a domain controller.
    8. In the features window, you do not have to select any additional features.
    9. Start the installation

+ Steps to promote the server to a domain controller after installation:
    1. After installation, there will be a clickable link "Promote this server to a domain controller"
    2. If you exited out of that wizard, on the top right, there will be a warning "Post-deployemnt configurations", you will find the "Promote this server to a domain controller" link there.
    3. If you are starting from scratch, select "Add a new forest" with a root domain name.
    4. For the forest functional level, select the oldest OS version you will run on a domain controller.
    5. Domain functional level should not be lower than the forest functional level.
    6. Set a DSRM password (This is a password to restore Active Directory)
    7. Select both GC and DNS options in the checkboxes.
    8. Choose next on the DNS delegation options window.
    9. Select next on the additional options window.
    10. Choose next on the paths window if you are ok with the default paths.
    11. You can export the deployement script with 'View Script'

+ Steps to verify sucessful Active Directory installation:
    1. Go to Server Manager > Tools > Event Viewer
    2. Go to Applications & Services Logs > Directory Service
    3. Check if there are any warnings. You can ignore DNS and 'None' category warnings.
    4. Go to Server Manager > Tools > DNS
        + Expand your DC > Forward Lookup Zone > msdcs.domain.machine > dc > tcp
        + Check for an LDAP and Kerberos SRV
    5. Go to Start > Run > \\domain.machineName
        + If 'NETLOGON' and 'SYSVOL' are available, the installation went okay.

+ How to install Active Directory through PowerShell:
    1. `Install-WindowsFeature ad-domain-services`
        + This is to install the AD services.
    2. `Get-WindowsFeature`
        + This is to ensure that the service got installed properly.
    3. `Install-ADDSDomainController -DomainName domain.name -SiteName default-first-site-name -InstallDns`
        + This is to promote the server to a DC. A password will be prompted.
    4. `Install-WindowsFeature rsat-role-tools`
        + This is to install AD tools.

+ Understanding Active Directory Schema
