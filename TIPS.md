## How-To

### List all workspaces (and display only their ID)
tecli workspace list --organization=terraform-cloud-pipeline | grep "ID" | grep "ws-" | awk '{ print $2}' | sed 's,",,g' | sed 's/,//g'