# Tips

This page collects task-oriented recipes that combine several TECLI commands. For the full command reference, see [COMMANDS.md](COMMANDS.md). For the most common single commands, see [TOP-COMMANDS.md](TOP-COMMANDS.md).

All examples read the organization and tokens from the active profile or the `TFC_*` environment variables. See [Configuration](README.md#configuration).

## Create a workspace linked to a VCS repository

1. List the OAuth tokens to find the token ID for your VCS connection:

```bash
tecli o-auth-token list
```

2. Create the workspace and pass the OAuth token ID and repository identifier:

```bash
tecli workspace create \
  --name terraform-dummy-1 \
  --vcs-repo-oauth-token-id ot-XXXXXXXX \
  --vcs-repo-identifier valter-silva-au/terraform-dummy
```

## Create a plan and apply

1. Find the workspace ID:

```bash
tecli workspace list
```

2. Create a run on the workspace:

```bash
tecli run create --workspace-id ws-XXXXXXXX --message "Plan and apply"
```

3. Apply the run by its ID:

```bash
tecli run apply --id run-XXXXXXXX --comment "Applying changes"
```

4. Read the apply by its ID:

```bash
tecli apply read --id apply-XXXXXXXX
```

Example output:

```json
{
  "ID": "apply-XXXXXXXX",
  "LogReadURL": "https://archivist.terraform.io/v1/object/...",
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
```

5. Read the apply logs:

```bash
tecli apply logs --id apply-XXXXXXXX
```

Example output:

```text
Terraform v0.14.5
Initializing plugins and modules...

Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

Outputs:

aws_region = "us-east-1"
```

## List workspace IDs only

Filter the `list` output down to the workspace IDs:

```bash
tecli workspace list | grep "ID" | grep "ws-" | awk '{ print $2 }' | sed 's,",,g' | sed 's/,//g'
```

## Set an OAuth token's private SSH key

Pass the key with a shell variable so the multi-line value is preserved:

```bash
private_ssh_key='-----BEGIN RSA PRIVATE KEY-----
... shortened for brevity ...
-----END RSA PRIVATE KEY-----'
tecli o-auth-token update --id ot-XXXXXXXX --private-ssh-key "${private_ssh_key}"
```

Example output:

```json
{
  "ID": "ot-XXXXXXXX",
  "UID": "",
  "CreatedAt": "2021-01-28T22:49:10.805Z",
  "HasSSHKey": true,
  "ServiceProviderUser": "valter-silva-au",
  "OAuthClient": {
    "ID": "oc-XXXXXXXX",
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
