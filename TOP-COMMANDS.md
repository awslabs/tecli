# Top commands

This is a quick-reference cheat sheet of the most common TECLI invocations. For the full command surface, see [COMMANDS.md](COMMANDS.md).

These examples assume you have a configured profile or the matching `TFC_*` environment variables set. The organization and tokens are read from the active profile or environment, not from a command-line flag. The token you need depends on the operation; most workspace and run operations use a [team API token](https://www.terraform.io/docs/cloud/users-teams-organizations/api-tokens.html#team-api-tokens).

## One-time setup

Run the interactive configure command, or export the environment variables.

```bash
tecli configure create
```

```bash
# Linux and macOS
export TFC_ORGANIZATION=your-org
export TFC_USER_TOKEN=your-user-token
export TFC_TEAM_TOKEN=your-team-token
export TFC_ORGANIZATION_TOKEN=your-org-token
```

```powershell
# Windows (PowerShell)
$Env:TFC_ORGANIZATION = "your-org"
$Env:TFC_USER_TOKEN = "your-user-token"
$Env:TFC_TEAM_TOKEN = "your-team-token"
$Env:TFC_ORGANIZATION_TOKEN = "your-org-token"
```

## Workspaces

List all workspaces in the organization on the active profile:

```bash
tecli workspace list
```

Find a workspace by name. This avoids paging through `list`:

```bash
tecli workspace find-by-name --name "${TFC_WORKSPACE_NAME}"
```

Create a workspace and allow destroy plans:

```bash
tecli workspace create \
  --name "${TFC_WORKSPACE_NAME}" \
  --allow-destroy-plan=true
```

Lock and unlock a workspace by ID:

```bash
tecli workspace lock --id "${WORKSPACE_ID}"
tecli workspace unlock --id "${WORKSPACE_ID}"
```

## Runs

Upload a configuration and create a run:

```bash
tecli configuration-version create --workspace-id "${WORKSPACE_ID}"
tecli configuration-version upload --url "${CV_UPLOAD_URL}" --path ./
tecli run create --workspace-id "${WORKSPACE_ID}" --message "${MESSAGE}"
```

Read a run's status:

```bash
tecli run read --id "${RUN_ID}"
```

Poll until the run leaves the pending state:

```bash
while true; do
  STATUS=$(tecli run read --id "${RUN_ID}" | jq -r ".Status")
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
tecli plan logs --id "${PLAN_ID}"
```

Create a destroy run:

```bash
tecli run create \
  --workspace-id "${WORKSPACE_ID}" \
  --message "${MESSAGE}" \
  --is-destroy=true
```

Discard one run, or discard everything queued on a workspace:

```bash
tecli run discard --id "${RUN_ID}"
tecli run discard-all --workspace-id "${WORKSPACE_ID}"
```

Apply a run and read the apply logs:

```bash
tecli run apply --id "${RUN_ID}" --comment "${COMMENT}"
tecli apply logs --id "${APPLY_ID}"
```

## Variables

Create a sensitive Terraform variable:

```bash
tecli variable create \
  --workspace-id "${WORKSPACE_ID}" \
  --key "${VARIABLE_KEY}" \
  --value "${VARIABLE_VALUE}" \
  --category terraform \
  --sensitive=true
```

Create a sensitive environment variable:

```bash
tecli variable create \
  --workspace-id "${WORKSPACE_ID}" \
  --key "${VARIABLE_KEY}" \
  --value "${VARIABLE_VALUE}" \
  --category env \
  --sensitive=true
```

Load AWS credentials as environment-category variables:

```bash
tecli variable create --workspace-id "${WORKSPACE_ID}" --key AWS_ACCESS_KEY_ID     --value "${AWS_ACCESS_KEY_ID}"     --category env --sensitive=true
tecli variable create --workspace-id "${WORKSPACE_ID}" --key AWS_SECRET_ACCESS_KEY --value "${AWS_SECRET_ACCESS_KEY}" --category env --sensitive=true
tecli variable create --workspace-id "${WORKSPACE_ID}" --key AWS_DEFAULT_REGION    --value "${AWS_DEFAULT_REGION}"    --category env --sensitive=true

# Add a session token if you use temporary credentials
tecli variable create --workspace-id "${WORKSPACE_ID}" --key AWS_SESSION_TOKEN --value "${AWS_SESSION_TOKEN}" --category env --sensitive=true
```

Update a variable by its key instead of by ID:

```bash
tecli variable update-by-key \
  --workspace-id "${WORKSPACE_ID}" \
  --key "${VARIABLE_KEY}" \
  --value "${VARIABLE_VALUE}"
```

Delete every variable on a workspace. This is irreversible:

```bash
tecli variable delete-all --workspace-id "${WORKSPACE_ID}"
```

## VCS and OAuth

List OAuth tokens to find the token ID for a workspace VCS connection:

```bash
tecli o-auth-token list
```

Set an OAuth token's private SSH key:

```bash
private_ssh_key="$(cat ~/.ssh/id_rsa)"
tecli o-auth-token update --id "${OAUTH_TOKEN_ID}" --private-ssh-key "${private_ssh_key}"
```

Manage the OAuth client:

```bash
tecli o-auth-client list
tecli o-auth-client read   --id "${OAUTH_CLIENT_ID}"
tecli o-auth-client delete --id "${OAUTH_CLIENT_ID}"
```
