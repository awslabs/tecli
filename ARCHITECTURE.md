# Architecture

A one-page tour of the TECLI codebase, current as of the June 2026
modernization. New contributors should read this before opening their
first PR; deeper details live in the package-level Go doc and the
[wiki](https://github.com/awslabs/tecli/wiki).

## High-level shape

TECLI is a thin command-line wrapper around
[`hashicorp/go-tfe`](https://github.com/hashicorp/go-tfe). The binary is
produced by [`main.go`](main.go), which only calls
`cmd.Execute()` from the `cobra/cmd` package.

The dependency direction is:

```
cobra/cmd  -->  cobra/controller  -->  cobra/aid  -->  helper/, box/
                                  -->  hashicorp/go-tfe
```

Layers below `controller` never import `cmd`; layers below `aid` never
import `controller`.

## Directory layout

| Path                | Role                                                                                                                                                                                                              |
| ------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `main.go`           | Process entry point. Calls `cmd.Execute()`.                                                                                                                                                                       |
| `cobra/cmd/`        | **Thin adapters.** One file per top-level command (`workspace`, `run`, `apply`, `plan`, `configuration-version`, `configure`, `o-auth-client`, `o-auth-token`, `ssh-key`, `variable`, `version`). Each file pulls a `*cobra.Command` from the controller package and registers it on `rootCmd`. No business logic here. |
| `cobra/controller/` | **Business logic.** Builds each `cobra.Command` (with `Use`/`Short`/`Long`/`Example` filled from `box/resources/manual/*.yaml`), wires `PreRunE`/`RunE`, validates flags via `helper.ValidateCmdArgs*`, and calls `go-tfe` to talk to Terraform Cloud/Enterprise. |
| `cobra/aid/`        | **Option builders + file I/O.** `SetXxxFlags(cmd)` registers per-command flags; `helper.go` and friends marshal flag values into `tfe.XxxOptions{}`, read/write the credentials file, and load viper config.       |
| `cobra/dao/`        | Data-access shims (currently small) for talking to the embedded resources.                                                                                                                                        |
| `cobra/model/`      | Plain structs. Most notably `CredentialProfile` used by the `configure` command.                                                                                                                                  |
| `cobra/view/`       | Output rendering helpers.                                                                                                                                                                                         |
| `helper/`           | **General utilities.** `cobra.go` (arg/flag validation), `directories.go`, `files.go`, `manual.go` (`GetManual` reads YAML manuals out of `box`), `ssh.go`, `strings.go`. No domain knowledge of Terraform Cloud. |
| `box/`              | **Embedded resources** (legacy `clencli` flow). `box.go` exposes the embedded blob; `gen.go` regenerates it. `resources/manual/*.yaml` defines each cobra command's `Use`/`Short`/`Long`/`Example`. `resources/VERSION` is the version string surfaced by `tecli version`. |
| `clencli/`          | Templates and assets consumed by `clencli` to render the README and screenshots: `readme.tmpl`, `readme.yaml`, `terminalizer/*.gif`, `logo.jpeg`.                                                                  |
| `habits/`           | Submodule (currently empty in tree) hosting shared Make targets included from `Makefile` (e.g. `go/build`, `go/fmt`, `go/install`).                                                                                |
| `examples/`         | End-user-facing usage examples (e.g. `examples/gitlab/` for GitLab CI).                                                                                                                                            |
| `tests/`            | Integration tests that hit Terraform Cloud — require `TFC_*` env vars or a configured profile to run. `tests/commands/` houses per-command test files.                                                            |
| `.github/workflows/`| CI: `build.yml` (per-push build), `publish.yml` (tag-driven release), `release.yml` (release-please).                                                                                                              |

## Adding a new command

1. Add a YAML manual at `box/resources/manual/<name>.yaml` describing
   `use`, `short`, `long`, `example`.
2. Add `cobra/controller/<name>.go` that:
   - Declares `var <name>ValidArgs = []string{ ... }`.
   - Builds a `*cobra.Command` from `helper.GetManual("<name>", validArgs)`.
   - Implements `<name>PreRun` (flag validation) and `<name>Run`
     (calls go-tfe).
3. Add `cobra/aid/<name>.go` with `SetXxxFlags(cmd *cobra.Command)`
   and option-builder helpers.
4. Register the command in `cobra/cmd/<name>.go` and add the
   `rootCmd.AddCommand(<name>Cmd)` call.
5. Run `go build ./...` and `go vet ./...`. If the command has tests,
   add them under `tests/commands/`.
6. Document the new command in `COMMANDS.md` and (if relevant)
   `TOP-COMMANDS.md`.

## Configuration & credentials

`tecli configure create` writes a YAML credentials file under
`~/.tecli/credentials.yaml` with one or more named profiles. Each
profile holds `organization`, `user-token`, `team-token`, and
`organization-token`. The same values can be supplied via
`TFC_ORGANIZATION`, `TFC_USER_TOKEN`, `TFC_TEAM_TOKEN`, and
`TFC_ORGANIZATION_TOKEN` environment variables (env wins).

Most commands accept `--profile`/`-p` (default: `default`) so a single
host can target multiple Terraform organizations.

## Build & release entry points

- Local build: `go build ./...` (requires Go 1.25+).
- Cross-compile: `make tecli/compile` (writes binaries to `dist/`).
- Release: see [`docs/RELEASING.md`](docs/RELEASING.md).
