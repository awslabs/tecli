## Commands
```
Command Line Interface for Terraform Enterprise/Cloud

Usage:
   [command]

Available Commands:
  apply                 An apply represents the results of applying a Terraform Run's execution plan.
  configuration-version A configuration version is a resource used to reference the uploaded configuration files.
  configure             Configures tecli settings
  help                  Help about any command
  o-auth-client         An OAuth Client represents the connection between an organization and a VCS provider.
  o-auth-token          The oauth-token object represents a VCS configuration which includes the OAuth connection and the associated OAuth token. This object is used when creating a workspace to identify which VCS connection to use.
  plan                  A plan represents the execution plan of a Run in a Terraform workspace.
  run                   A run performs a plan and apply, using a configuration version and the workspace’s current variables.
  ssh-key               The ssh-key object represents an SSH key which includes a name and the SSH private key. An organization can have multiple SSH keys available.
  variable              Operations on variables.
  version               Displays the version of tecli and all installed plugins
  workspace             Workspaces represent running infrastructure managed by Terraform.

Flags:
  -h, --help             help for this command
  -p, --profile string   Use a specific profile from your credentials and configurations file. (default "default")

Use " [command] --help" for more information about a command.
```
