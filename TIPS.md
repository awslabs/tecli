## How-To

### Create a workspace linked with a VCS repository
```
tecli oauth list --organization terraform-cloud-pipeline
tecli workspace create --vcs-repo-oauth-token-id=ot-xEpT34KFFtKXXXXX --vcs-repo-identifier=valter-silva-au/terraform-dummy --organization=terraform-cloud-pipeline --name terraform-dummy-1
```

### Create a plan and apply 
```
# Get the workspace ID you want to manage
tecli workspace list --organization=terraform-cloud-pipeline

# Create a run (plan) on the workspace and get the apply ID
tecli run create --workspace-id ws-eMVVNYy4xhyUMtWq

# Apply the run by providing the ID
tecli run apply --id=run-9NPyVKuyoDAdU925

# Read the apply by providing the apply ID
tecli apply read --id=apply-J7VA5BDsXEEkxccU

{
  "ID": "apply-J7VA5BDsXEEkxccU",
  "LogReadURL": "https://archivist.terraform.io/v1/object/dmF1bHQ6djE6ZTdScXFZSS9nZjlYdE85V0I0R1RDQW5PbXdxd21qdVBwdTlEeXdMWnovNnY3SjZKa25GYXFUcDBpbDl5MGRHbHZzak41eFJIbjQraDVuSzV0WkFaSVQ2YkZOaWpjSi84TWNJT2lPL1N0VTJSVU80U0ZVSUp1elE0cEI4Zm9ZQ0FESGRlWkVUM2h1cXNSNnQzR0ZqRFVVRHYyRlFoeGxibCtSQ1F5V0VlWlZwMTF3cUxTUDA4ckNsUG4xWW5wdGZ5K0o4M3AvUnJ4cXo4LzI3TTVta1czdWFXRFlkSHFMTC8zRHFpS0VGbG9LNmh3a2NpdmtuMnc1QWJ3UGVBRkNSeXFkUStyMVowWGxxbmh1NzZCcXFDU1drVzR1QTBGb1pmd0xhOWttRkVvcGlnZCs4THJmVkV1Q3JzRzlrOTVsMmxrVktDNnRXalBuZHRLcUhLM3dqMWhUdXFWWHRQYzFNYm1PY2l0L1JIeWVCeTZ4K3FraHlxbEI3U2d6UEtmSGpS",
  "ResourceAdditions": 0,
  "ResourceChanges": 0,
  "ResourceDestructions": 0,
  "Status": "finished",
  "StatusTimestamps": {
    "canceled-at": "0001-01-01T00:00:00Z",
    "errored-at": "0001-01-01T00:00:00Z",
    "finished-at": "2021-01-28T23:39:22Z",
    "force-canceled-at": "0001-01-01T00:00:00Z",
    "queued-at": "0001-01-01T00:00:00Z",
    "started-at": "2021-01-28T23:39:17Z"
  }
}

# Read the apply logs
tecli apply logs --id=apply-J7VA5BDsXEEkxccU

Terraform v0.14.5
Initializing plugins and modules...

Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

Outputs:

aws_region = "us-east-1"

```

### List all workspaces (and display only their ID)
```
tecli workspace list --organization=terraform-cloud-pipeline | grep "ID" | grep "ws-" | awk '{ print $2}' | sed 's,",,g' | sed 's/,//g'
```

### How to update o-auth-token
```
$ private_ssh_key='-----BEGIN RSA PRIVATE KEY-----
.. shorten for brevity
> -----END RSA PRIVATE KEY-----
> '
$ tecli o-auth-token update --id=ot-C7MnH5Qx8aQXXXXX --private-ssh-key "${private_ssh_key}"
{
  "ID": "ot-C7MnH5Qx8aQXXXXX",
  "UID": "",
  "CreatedAt": "2021-01-28T22:49:10.805Z",
  "HasSSHKey": true,
  "ServiceProviderUser": "valter-silva-au",
  "OAuthClient": {
    "ID": "oc-psMU7aX65hiXXXXX",
    "APIURL": "",
    "CallbackURL": "",
    "ConnectPath": "",
    "CreatedAt": "0001-01-01T00:00:00Z",
    "HTTPURL": "",
    "Key": "",
    "RSAPublicKey": "",
    "ServiceProvider": "",
    "ServiceProviderName": "",
    "Organization": null,
    "OAuthTokens": null
  }
}
```
