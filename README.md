# AWS Profile Manager

## Description
Choose your default AWS profile from your list of existing profiles.

The access key and secret key of the selected profile will be set as the keys of the default profile

---

## Requisites

Go version 1.20 or greater.

https://go.dev/doc/install

---

## Installation

> [!WARNING]  
> Create backup of your credentials file!

Windows:
```powershell
Copy-Item "C:\Users\$Env:UserName\.aws\credentials" -Destination "C:\Users\${Env:UserName}\.aws\credentials.bak"
```

Linux:
```bash
cp ~/.aws/credentials ~/.aws/credentials.bak
```

After the backup, clone the repository and run install command:

```bash
git clone https://github.com/delanovictor/aws-profile-manager.git

cd aws-profile-manager

go install

# Linux only
export PATH=$PATH:~/go/bin/

#Suggested alias
alias awsp=aws-profile
```
--- 
## Example

```
aws-profile ls

//   [profile1]
// =>[profile2]
//   [profile3]

aws-profile set profile1

// Selected [profile1]

aws-profile ls

// =>[profile1]
//   [profile2]
//   [profile3]

```

