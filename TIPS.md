## How-To

### Create a workspace linked with a VCS repository

```
tecli oauth list --organization terraform-cloud-pipeline
tecli workspace create --vcs-repo-oauth-token-id=ot-XXX --vcs-repo-identifier=valter-silva-au/terraform-dummy --organization=terraform-cloud-pipeline --name terraform-dummy-1
```

### Create a plan and apply

```
# Get the workspace ID you want to manage
tecli workspace list --organization=terraform-cloud-pipeline

# Create a run (plan) on the workspace and get the apply ID
tecli run create --workspace-id ws-XXX

# Apply the run by providing the ID
tecli run apply --id=run-XXX

# Read the apply by providing the apply ID
tecli apply read --id=apply-XXX

{
  "ID": "apply-XXX",
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

# Read the apply logs
tecli apply logs --id=apply-XXX

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
