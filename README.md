# AWS Profile Manager

## Description
Choose your default AWS profile from your list of existing profiles.

The access key and secrey key of the selected profile will be set as the keys of the default profile

---

## Installation

```
# IMPORTANT: Create backup of credentials file

cp ~/.aws/credentials ~/.aws/credentials.bak

git clone https://github.com/delanovictor/aws-profile-manager.git

cd aws-profile-manager

go install

# Add to .bashrc
export PATH=$PATH:~/go/bin/
```
--- 
## Example


```
aws-profile ls

//   [profile1]
// =>[profile2]
//   [profile3]

aws-profile select profile1

// Selected [profile1]

aws-profile ls

// =>[profile1]
//   [profile2]
//   [profile3]

```

