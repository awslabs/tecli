# Commands

This is the full command reference for TECLI. It is verified against the `*ValidArgs` slices and flag definitions in [`cobra/controller/`](cobra/controller/) and [`cobra/aid/`](cobra/aid/).

Run `tecli --help` for the live top-level list and `tecli <command> --help` for flag-level help. If this file and `tecli --help` disagree, open an issue or a pull request.

## Conventions

- A command has the form `tecli <command> <argument> [flags]`.
- `<command>` is a resource. `<argument>` is the operation, such as `list` or `create`.
- The organization and API tokens are not command-line flags. TECLI reads them from the `TFC_*` environment variables or the active profile. See [Configuration](README.md#configuration).
- Identifiers such as `--id` and `--workspace-id` are Terraform Cloud resource IDs (for example, `ws-XXXXXXXX`), not names.

## Persistent flags

These flags are available on every command.

| Flag              | Default   | Description                                        |
| ----------------- | --------- | -------------------------------------------------- |
| `-p`, `--profile` | `default` | Selects a named profile from the credentials file. |
| `-h`, `--help`    |           | Prints help for the command.                       |

## `tecli configure`

Manages the TECLI credentials file. `configure` reads and writes the credentials file only. It does not use environment variables.

Arguments: `list`, `create`, `read`, `update`, `delete`.

| Flag                   | Default       | Description                                                      |
| ---------------------- | ------------- | ---------------------------------------------------------------- |
| `--mode`               | `interactive` | `interactive` prompts for values; `non-interactive` reads flags. |
| `--new-name`           |               | New profile name for `update`.                                   |
| `--description`        |               | Profile description.                                             |
| `--user-token`         |               | User API token (non-interactive mode).                           |
| `--team-token`         |               | Team API token (non-interactive mode).                           |
| `--organization-token` |               | Organization API token (non-interactive mode).                   |

```bash
# Create the default profile interactively
tecli configure create

# Create a named profile non-interactively
tecli configure create --profile cicd --mode non-interactive \
  --team-token "${TFC_TEAM_TOKEN}"

# Read and list profiles
tecli configure read --profile cicd
tecli configure list

# Delete a profile
tecli configure delete --profile cicd
```

## `tecli workspace`

Manages workspaces. Viewing a workspace requires permission to read runs. Changing settings and force-unlocking require admin access. Locking and unlocking require lock and unlock permission.

Arguments: `list`, `create`, `read`, `read-by-id`, `update`, `update-by-id`, `delete`, `delete-by-id`, `find-by-name`, `lock`, `unlock`, `force-unlock`, `assign-ssh-key`, `unassign-ssh-key`, `remove-vcs-connection`, `remove-vcs-connection-by-id`.

Name-based arguments (`create`, `read`, `update`, `delete`, `find-by-name`, `remove-vcs-connection`) require `--name`. ID-based arguments (`read-by-id`, `update-by-id`, `delete-by-id`, `remove-vcs-connection-by-id`, `lock`, `unlock`, `force-unlock`, `assign-ssh-key`, `unassign-ssh-key`) require `--id`.

| Flag                            | Type        | Description                                            |
| ------------------------------- | ----------- | ------------------------------------------------------ |
| `--id`                          | string      | Workspace ID (`ws-XXXXXXXX`).                          |
| `--name`                        | string      | Workspace name.                                        |
| `--new-name`                    | string      | New name for `update`.                                 |
| `--search`                      | string      | Search filter for `list`.                              |
| `--include`                     | string      | Related resources to include in `list`.                |
| `--agent-pool-id`               | string      | Agent pool ID.                                         |
| `--allow-destroy-plan`          | bool        | Allow destroy plans.                                   |
| `--auto-apply`                  | bool        | Apply runs automatically.                              |
| `--execution-mode`              | string      | Execution mode (`remote`, `local`, `agent`).           |
| `--file-triggers-enabled`       | bool        | Trigger runs on file changes in `--trigger-prefixes`.  |
| `--migration-environment`       | string      | Legacy migration environment.                          |
| `--queue-all-runs`              | bool        | Queue all runs on creation.                            |
| `--speculative-enabled`         | bool        | Allow speculative plans.                               |
| `--terraform-version`           | string      | Terraform version for the workspace.                   |
| `--trigger-prefixes`            | stringArray | Path prefixes that trigger runs.                       |
| `--working-directory`           | string      | Working directory for Terraform.                       |
| `--reason`                      | string      | Reason for `lock`.                                     |
| `--ssh-key-id`                  | string      | SSH key ID for `assign-ssh-key`.                       |
| `--vcs-repo-branch`             | string      | VCS branch.                                            |
| `--vcs-repo-identifier`         | string      | VCS repository identifier (`org/repo`).                |
| `--vcs-repo-ingress-submodules` | bool        | Fetch submodules when cloning.                         |
| `--vcs-repo-oauth-token-id`     | string      | OAuth token ID for the VCS connection (`ot-XXXXXXXX`). |

```bash
# List workspaces in the organization on the active profile
tecli workspace list

# Find a workspace by name
tecli workspace find-by-name --name your-workspace

# Create a workspace that allows destroy plans
tecli workspace create --name your-workspace --allow-destroy-plan=true

# Create a workspace connected to a VCS repository
tecli workspace create \
  --name your-workspace \
  --vcs-repo-oauth-token-id ot-XXXXXXXX \
  --vcs-repo-identifier org/repo

# Lock and unlock a workspace by ID
tecli workspace lock --id ws-XXXXXXXX
tecli workspace unlock --id ws-XXXXXXXX
```

## `tecli run`

Manages runs. A run performs a plan and apply using a configuration version and the workspace's current variables.

Arguments: `list`, `create`, `read`, `read-with-options`, `apply`, `cancel`, `cancel-all`, `force-cancel`, `force-cancel-all`, `discard`, `discard-all`.

`list`, `create`, `cancel-all`, `force-cancel-all`, and `discard-all` operate on a workspace and require `--workspace-id`. `read`, `read-with-options`, `apply`, `cancel`, `force-cancel`, and `discard` operate on a single run and require `--id`.

| Flag                         | Type        | Description                                                   |
| ---------------------------- | ----------- | ------------------------------------------------------------- |
| `--id`                       | string      | Run ID (`run-XXXXXXXX`).                                      |
| `--workspace-id`             | string      | Workspace ID (`ws-XXXXXXXX`).                                 |
| `--configuration-version-id` | string      | Configuration version ID to run against.                      |
| `--message`                  | string      | Message associated with the run (used by `create`).           |
| `--comment`                  | string      | Comment for `apply`, `cancel`, `force-cancel`, and `discard`. |
| `--is-destroy`               | bool        | Create a destroy run.                                         |
| `--target-addrs`             | stringArray | Resource addresses to target.                                 |
| `--include`                  | string      | Related resources to include in the read.                     |

```bash
# Create a run on a workspace
tecli run create --workspace-id ws-XXXXXXXX --message "Initial run"

# Create a destroy run
tecli run create --workspace-id ws-XXXXXXXX --message "Tear down" --is-destroy=true

# Read a run
tecli run read --id run-XXXXXXXX

# Apply a run with a comment
tecli run apply --id run-XXXXXXXX --comment "Applying changes"

# Discard one run, or every run queued on a workspace
tecli run discard --id run-XXXXXXXX
tecli run discard-all --workspace-id ws-XXXXXXXX
```

## `tecli plan`

Reads plans and plan logs. A plan is the execution plan of a run.

Arguments: `read`, `logs`. Both require `--id` (the plan ID).

| Flag   | Type   | Description |
| ------ | ------ | ----------- |
| `--id` | string | Plan ID.    |

```bash
tecli plan read --id plan-XXXXXXXX
tecli plan logs --id plan-XXXXXXXX
```

## `tecli apply`

Reads applies and apply logs. An apply represents the results of applying a run's execution plan.

Arguments: `read`, `logs`. Both require `--id` (the apply ID).

| Flag   | Type   | Description |
| ------ | ------ | ----------- |
| `--id` | string | Apply ID.   |

```bash
tecli apply read --id apply-XXXXXXXX
tecli apply logs --id apply-XXXXXXXX
```

## `tecli configuration-version`

Manages configuration versions. A configuration version references the uploaded configuration files used by a run.

Arguments: `list`, `create`, `read`, `upload`. `list` and `create` require `--workspace-id`. `read` requires `--id`. `upload` requires `--url` and `--path`.

| Flag                | Type   | Description                                               |
| ------------------- | ------ | --------------------------------------------------------- |
| `--id`              | string | Configuration version ID (`cv-XXXXXXXX`).                 |
| `--workspace-id`    | string | Workspace ID (`ws-XXXXXXXX`).                             |
| `--auto-queue-runs` | bool   | Queue a run automatically after upload.                   |
| `--speculative`     | bool   | Mark the configuration version as speculative.            |
| `--url`             | string | Upload URL returned by `create` (used by `upload`).       |
| `--path`            | string | Local path to the configuration files (used by `upload`). |

```bash
# Create a configuration version and capture the upload URL
tecli configuration-version create --workspace-id ws-XXXXXXXX

# Upload configuration files to that URL
tecli configuration-version upload --url https://archivist.terraform.io/... --path ./

# Read a configuration version
tecli configuration-version read --id cv-XXXXXXXX
```

## `tecli variable`

Manages Terraform and environment variables on a workspace.

Arguments: `list`, `create`, `read`, `update`, `update-by-key`, `delete`, `delete-all`. `list`, `create`, and `delete-all` operate on a workspace and require `--workspace-id`. `read`, `update`, and `delete` require `--id`. `update-by-key` matches a variable by `--key`.

| Flag             | Type   | Description                     |
| ---------------- | ------ | ------------------------------- |
| `--id`           | string | Variable ID (`var-XXXXXXXX`).   |
| `--workspace-id` | string | Workspace ID (`ws-XXXXXXXX`).   |
| `--key`          | string | Variable key.                   |
| `--value`        | string | Variable value.                 |
| `--description`  | string | Variable description.           |
| `--category`     | string | `terraform` or `env`.           |
| `--hcl`          | bool   | Parse the value as HCL.         |
| `--sensitive`    | bool   | Mark the variable as sensitive. |

```bash
# Create a sensitive Terraform variable
tecli variable create \
  --workspace-id ws-XXXXXXXX \
  --key your-key \
  --value your-value \
  --category terraform \
  --sensitive=true

# Create a sensitive environment variable
tecli variable create \
  --workspace-id ws-XXXXXXXX \
  --key AWS_SECRET_ACCESS_KEY \
  --value your-secret-key \
  --category env \
  --sensitive=true

# Update a variable by key
tecli variable update-by-key --workspace-id ws-XXXXXXXX --key your-key --value new-value

# Delete every variable on a workspace
tecli variable delete-all --workspace-id ws-XXXXXXXX
```

## `tecli ssh-key`

Manages SSH keys. SSH keys are used by VCS integrations and by workspaces that clone modules from a Git server. The list and read operations return metadata only; Terraform Cloud never returns the private key text.

Arguments: `list`, `create`, `read`, `update`, `delete`. `read`, `update`, and `delete` require `--id`.

| Flag      | Type   | Description                                          |
| --------- | ------ | ---------------------------------------------------- |
| `--id`    | string | SSH key ID. Required for `read`, `update`, `delete`. |
| `--name`  | string | Name to identify the SSH key.                        |
| `--value` | string | Content of the SSH private key.                      |

```bash
tecli ssh-key list
tecli ssh-key read --id sshkey-XXXXXXXX
tecli ssh-key delete --id sshkey-XXXXXXXX
```

## `tecli o-auth-client`

Manages OAuth clients. An OAuth client represents the connection between an organization and a VCS provider.

Arguments: `list`, `create`, `read`, `delete`. `read` and `delete` require `--id`.

| Flag                 | Type   | Description                        |
| -------------------- | ------ | ---------------------------------- |
| `--id`               | string | OAuth client ID (`oc-XXXXXXXX`).   |
| `--api-url`          | string | VCS provider API URL.              |
| `--http-url`         | string | VCS provider HTTP URL.             |
| `--o-auth-token`     | string | OAuth token from the VCS provider. |
| `--private-key`      | string | Private key for the connection.    |
| `--service-provider` | string | VCS service provider type.         |

```bash
tecli o-auth-client list
tecli o-auth-client read --id oc-XXXXXXXX
tecli o-auth-client delete --id oc-XXXXXXXX
```

## `tecli o-auth-token`

Manages OAuth tokens. An OAuth token is the VCS configuration used when you create a workspace connected to a repository.

Arguments: `list`, `read`, `update`, `delete`. `read`, `update`, and `delete` require `--id`.

| Flag                | Type   | Description                                    |
| ------------------- | ------ | ---------------------------------------------- |
| `--id`              | string | OAuth token ID (`ot-XXXXXXXX`).                |
| `--private-ssh-key` | string | Private SSH key used for git clone operations. |

```bash
# List OAuth tokens to find the ID for a VCS connection
tecli o-auth-token list

# Set the private SSH key on an OAuth token
private_ssh_key="$(cat ~/.ssh/id_rsa)"
tecli o-auth-token update --id ot-XXXXXXXX --private-ssh-key "${private_ssh_key}"
```

## `tecli version`

Prints the TECLI version. The version string is read from `box/resources/VERSION`. This command takes no arguments.

```bash
tecli version
```

## Related references

- Copy-paste recipes for common tasks: [TOP-COMMANDS.md](TOP-COMMANDS.md).
- Where these commands live in the codebase: [ARCHITECTURE.md](ARCHITECTURE.md).
