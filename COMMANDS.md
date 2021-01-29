## Commands
```
Terraform Enterprise/Cloud Command Line Interface

Usage:
   [command]

Available Commands:
  apply                 An apply represents the results of applying a Terraform Run's execution plan.
  configuration-version 
  configure             Configures tecli global settings
  help                  Help about any command
  o-auth-client         An OAuth Client represents the connection between an organization and a VCS provider.
  o-auth-token          The oauth-token object represents a VCS configuration which includes the OAuth connection and the associated OAuth token. This object is used when creating a workspace to identify which VCS connection to use.
  plan                  A plan represents the execution plan of a Run in a Terraform workspace.
  run                   
  ssh-key               
  version               Displays the version of tecli and all installed plugins
  workspace             Manage Terraform Cloud workspaces

Flags:
  -h, --help                   help for this command
  -l, --log string             Enable or disable logs (found at $HOME/.tecli/logs.json). Log outputs will be shown on default output. (default "disable")
      --log-file-path string   Log file path. (default "/Users/valterh/.tecli/logs.json")
  -o, --organization string    Terraform Cloud Organization name
  -p, --profile string         Use a specific profile from your credentials and configurations file. (default "default")
  -v, --verbosity string       Valid log level:panic,fatal,error,warn,info,debug,trace). (default "error")

Use " [command] --help" for more information about a command.
```
