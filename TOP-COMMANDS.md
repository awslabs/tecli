# Top Commands

Hand-curated cheat sheet of the most common TECLI invocations. The
full command surface lives in [`COMMANDS.md`](COMMANDS.md).

All examples assume you have a configured profile or the appropriate
[`TFC_*` environment variables](https://github.com/awslabs/tecli/wiki/Environment-Variables)
set. The relevant tokens depend on the operation; most workspace and
run operations need a [team API token](https://www.terraform.io/docs/cloud/users-teams-organizations/api-tokens.html#team-api-tokens).

## One-time setup

Either run `tecli configure create` and follow the prompts, or export:

```bash
# Linux / macOS
export TFC_ORGANIZATION=your-org
export TFC_USER_TOKEN=your-user-token
export TFC_TEAM_TOKEN=your-team-token
export TFC_ORGANIZATION_TOKEN=your-org-token

# Windows (PowerShell)
$Env:TFC_ORGANIZATION="your-org"
$Env:TFC_USER_TOKEN="your-user-token"
$Env:TFC_TEAM_TOKEN="your-team-token"
$Env:TFC_ORGANIZATION_TOKEN="your-org-token"
```

## Workspace cookbook

List all workspaces in an organization:

```bash
tecli workspace list -o "${TFC_ORGANIZATION}" -p "${PROFILE}"
```

Find a workspace by name (avoids paging through `list`):

```bash
tecli workspace find-by-name \
  --organization="${TFC_ORGANIZATION}" \
  --name="${TFC_WORKSPACE_NAME}"
```

Create a workspace and allow destroy plans:

```bash
tecli workspace create \
  --organization="${TFC_ORGANIZATION}" \
  --name="${TFC_WORKSPACE_NAME}" \
  --allow-destroy-plan=true
```

Lock / unlock a workspace:

```bash
tecli workspace lock   --id="${WORKSPACE_ID}"
tecli workspace unlock --id="${WORKSPACE_ID}"
```

## Run cookbook

Upload a configuration and create a run:

```bash
tecli configuration-version create --workspace-id="${WORKSPACE_ID}"
tecli configuration-version upload  --url="${CV_UPLOAD_URL}" --path=./
tecli run create --workspace-id="${WORKSPACE_ID}" --comment="${COMMENT}"
```

Check run status:

```bash
tecli run read --id="${RUN_ID}"
```

Poll until the run finishes (bash one-liner):

```bash
while true; do
  STATUS=$(tecli run read --id="${RUN_ID}" | jq -r ".Status")
  if [ "${STATUS}" != "pending" ]; then
    break
  else
    echo "RUN STATUS: ${STATUS} — sleeping 5s"
    sleep 5
  fi
done
```

Show plan logs:

```bash
tecli plan logs --id="${PLAN_ID}"
```

Create a destroy run:

```bash
tecli run create \
  --workspace-id="${WORKSPACE_ID}" \
  --comment="${COMMENT}" \
  --is-destroy=true
```

Discard a single run, or discard everything queued on a workspace:

```bash
tecli run discard     --id="${RUN_ID}"
tecli run discard-all --workspace-id="${WORKSPACE_ID}"
```

Apply a run and stream the apply logs:

```bash
tecli run apply --id="${RUN_ID}" --comment="${COMMENT}"
tecli apply logs --id="${APPLY_ID}"
```

## Variable cookbook

Create a sensitive Terraform variable:

```bash
tecli variable create \
  --key="${VARIABLE_KEY}" \
  --value="${VARIABLE_VALUE}" \
  --workspace-id="${WORKSPACE_ID}" \
  --category=terraform \
  --sensitive=true
```

Create a sensitive environment variable:

```bash
tecli variable create \
  --key="${VARIABLE_KEY}" \
  --value="${VARIABLE_VALUE}" \
  --workspace-id="${WORKSPACE_ID}" \
  --category=env \
  --sensitive=true
```

Bulk-load AWS credentials as env-category variables:

```bash
tecli variable create --key=AWS_ACCESS_KEY_ID     --value="${AWS_ACCESS_KEY_ID}"     --workspace-id="${WORKSPACE_ID}" --category=env --sensitive=true
tecli variable create --key=AWS_SECRET_ACCESS_KEY --value="${AWS_SECRET_ACCESS_KEY}" --workspace-id="${WORKSPACE_ID}" --category=env --sensitive=true
tecli variable create --key=AWS_DEFAULT_REGION    --value="${AWS_DEFAULT_REGION}"    --workspace-id="${WORKSPACE_ID}" --category=env --sensitive=true

# If you also need a session token:
tecli variable create --key=AWS_SESSION_TOKEN --value="${AWS_SESSION_TOKEN}" --workspace-id="${WORKSPACE_ID}" --category=env --sensitive=true
```

Update an existing variable by its key (instead of by ID):

```bash
tecli variable update-by-key \
  --key="${VARIABLE_KEY}" \
  --value="${VARIABLE_VALUE}" \
  --workspace-id="${WORKSPACE_ID}"
```

Wipe every variable on a workspace (irreversible):

```bash
tecli variable delete-all --workspace-id="${WORKSPACE_ID}"
```

## VCS / OAuth cookbook

List OAuth tokens (used to wire workspaces to a VCS repo):

```bash
tecli o-auth-token list --organization="${TFC_ORGANIZATION}"
```

Update an OAuth token's private SSH key (heredoc-friendly):

```bash
private_ssh_key="$(cat ~/.ssh/id_rsa)"
tecli o-auth-token update --id="${OAUTH_TOKEN_ID}" --private-ssh-key "${private_ssh_key}"
```

Manage the OAuth client itself:

```bash
tecli o-auth-client list   --organization="${TFC_ORGANIZATION}"
tecli o-auth-client read   --id="${OAUTH_CLIENT_ID}"
tecli o-auth-client delete --id="${OAUTH_CLIENT_ID}"
```
