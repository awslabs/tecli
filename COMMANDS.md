# Commands

This file is the canonical, hand-maintained map of every TECLI
subcommand. Each section corresponds to one `*ValidArgs` slice in
[`cobra/controller/`](cobra/controller/), so `grep cobra.Command{
cobra/controller/` is the source of truth if these ever drift.

`tecli --help` prints the live top-level list. If you spot a
divergence between this file and `tecli --help`, please open an issue
or PR.

## Top-level commands

```text
Command Line Interface for Terraform Enterprise/Cloud

Usage:
  tecli [command]

Available Commands:
  apply                 An apply represents the results of applying a Terraform Run's execution plan.
  configuration-version A configuration version is a resource used to reference the uploaded configuration files.
  configure             Configures tecli settings
  help                  Help about any command
  o-auth-client         An OAuth Client represents the connection between an organization and a VCS provider.
  o-auth-token          The oauth-token object represents a VCS configuration which includes the OAuth connection and the associated OAuth token. This object is used when creating a workspace to identify which VCS connection to use.
  plan                  A plan represents the execution plan of a Run in a Terraform workspace.
  run                   A run performs a plan and apply, using a configuration version and the workspace's current variables.
  ssh-key               The ssh-key object represents an SSH key which includes a name and the SSH private key. An organization can have multiple SSH keys available.
  variable              Operations on variables.
  version               Displays the version of tecli and all installed plugins
  workspace             Workspaces represent running infrastructure managed by Terraform.

Flags:
  -h, --help             help for this command
  -p, --profile string   Use a specific profile from your credentials and configurations file. (default "default")

Use "tecli [command] --help" for more information about a command.
```

## Subcommands

Use `tecli <command> --help` to see flag-level help. The lists below
are extracted from the `*ValidArgs` definitions in
`cobra/controller/`.

### `tecli apply`

`read`, `logs`

### `tecli configuration-version`

`list`, `create`, `read`, `upload`

### `tecli configure`

`list`, `create`, `read`, `update`, `delete`

### `tecli o-auth-client`

`list`, `create`, `read`, `delete`

### `tecli o-auth-token`

`list`, `read`, `update`, `delete`

### `tecli plan`

`read`, `logs`

### `tecli run`

`list`, `create`, `read`, `read-with-options`, `apply`, `cancel`,
`cancel-all`, `force-cancel`, `force-cancel-all`, `discard`,
`discard-all`

### `tecli ssh-key`

`list`, `create`, `read`, `update`, `delete`

### `tecli variable`

`list`, `create`, `read`, `update`, `update-by-key`, `delete`,
`delete-all`

### `tecli version`

No subcommands. Prints the version stored in
`box/resources/VERSION`.

### `tecli workspace`

`list`, `create`, `read`, `read-by-id`, `update`, `update-by-id`,
`delete`, `delete-by-id`, `find-by-name`, `lock`, `unlock`,
`force-unlock`, `assign-ssh-key`, `unassign-ssh-key`,
`remove-vcs-connection`, `remove-vcs-connection-by-id`

## Common flags

- `-p, --profile string` (persistent, default `default`) — selects a
  named profile from `~/.tecli/credentials.yaml`.
- `-h, --help` — per-command help. Always available.

## Cross-references

- One-liner recipes for the most common workflows live in
  [`TOP-COMMANDS.md`](TOP-COMMANDS.md).
- Architectural background (where these commands live in the codebase)
  is in [`ARCHITECTURE.md`](ARCHITECTURE.md).
