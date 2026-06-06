# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and the project loosely follows [Semantic Versioning](https://semver.org/spec/v2.0.0.html)
(pre-1.0 releases use the `0.x-alpha` convention).

## [Unreleased] - Modernization 2026-06

Toolchain modernization. Tracked in PR [#27](https://github.com/awslabs/tecli/pull/27)
and follow-up PRs branched from `chore/modernize-toolchain`.

### Changed

- Bumped `go.mod` directive to **Go 1.25**.
- Migrated `github.com/hashicorp/go-tfe` from `v0.15.0` to `v1.108.0`.
  Required adapting controller code in `cobra/controller/` to the new
  context-aware client API and updated Go-TFE structs.
- Bumped `github.com/spf13/cobra` to `v1.8.1`.
- Bumped `github.com/spf13/viper` to `v1.19.0`.
- Bumped `github.com/sirupsen/logrus` to `v1.9.3`.
- Bumped `github.com/stretchr/testify` to `v1.11.1`.
- Bumped `github.com/spf13/afero` to `v1.15.0`.
- Bumped `github.com/hashicorp/go-slug` to `v0.16.3`.
- Bumped `github.com/hashicorp/go-retryablehttp` to `v0.7.7`.
- Bumped `golang.org/x/crypto` to `v0.45.0`.

### Documentation

- Rebuilt `CHANGELOG.md` (this file) to follow Keep a Changelog.
- Refreshed `README.md` with current install/build instructions and
  a History section that points at the modernization PRs.
- Aligned `COMMANDS.md` and `TOP-COMMANDS.md` with the actual cobra
  subcommands shipped today; flagged divergences and removed wiki/Markdown
  rot.
- Updated `CONTRIBUTING.md` for the Go 1.25 toolchain, added a
  "How to run tests locally" section, and documented the modernization
  branch / PR flow.
- Added [`ARCHITECTURE.md`](ARCHITECTURE.md): one-page summary of the
  directory layout (`cobra/cmd/`, `cobra/controller/`, `cobra/aid/`,
  `helper/`, `box/`, `clencli/`, `habits/`, `examples/`, `tests/`).
- Added [`docs/RELEASING.md`](docs/RELEASING.md): release-please workflow
  and manual fallback steps.

### Removed

- Dropped unused indirect dependencies pinned by stale `// indirect` lines
  in `go.mod` (handled automatically by `go mod tidy`).

## [0.4.2-alpha] - 2021-2025

Last shipped pre-modernization line of releases. Changes from this era,
summarized from `git log` (no formal changelog was maintained at the
time):

### Added

- `tecli configure` family for managing credential profiles.
- `tecli workspace` commands: `list`, `create`, `read`, `read-by-id`,
  `update`, `update-by-id`, `delete`, `delete-by-id`, `find-by-name`,
  `lock`, `unlock`, `force-unlock`, `assign-ssh-key`,
  `unassign-ssh-key`, `remove-vcs-connection`,
  `remove-vcs-connection-by-id`.
- `tecli run` commands: `list`, `create`, `read`, `read-with-options`,
  `apply`, `cancel`, `cancel-all`, `force-cancel`, `force-cancel-all`,
  `discard`, `discard-all`.
- `tecli plan` commands: `read`, `logs`.
- `tecli apply` commands: `read`, `logs`.
- `tecli configuration-version` commands: `list`, `create`, `read`,
  `upload`.
- `tecli variable` commands: `list`, `create`, `read`, `update`,
  `update-by-key`, `delete`, `delete-all`.
- `tecli o-auth-client` and `tecli o-auth-token` for VCS provider
  configuration.
- `tecli ssh-key` for managing SSH keys against private modules.
- `TFC_ORGANIZATION`, `TFC_USER_TOKEN`, `TFC_TEAM_TOKEN`, and
  `TFC_ORGANIZATION_TOKEN` environment-variable support.
- DevContainer configuration for VS Code under `.devcontainer/`.
- Pre-commit configuration (`.pre-commit-config.yaml`) and
  `goimports` integration.
- `release-please` workflow scaffolding.
- GitLab CI example under `examples/gitlab/`.
- Logo, terminalizer GIFs, and `clencli` README templates under
  `clencli/`.

### Changed

- `tecli configure create` reads `TFC_ORGANIZATION` from the
  credentials file/env when present.
- Standard output unified on `fmt` instead of cobra's command writers.
- Cobra root command logic moved into the `controller` package for
  testability.
- `GetManual` rewritten to drive `Use`/`Short`/`Long`/`Example` from
  YAML files in `box/resources/manual/`.
- Switched from `spf13/cobra/cobra` to standalone `cobra-cli`
  (PR [#14](https://github.com/awslabs/tecli/pull/14)).
- Silenced cobra's automatic usage dump on errors.

### Fixed

- `configure create` bugs around profile resolution.
- Workspace access-level requirements for commands that did not need
  full admin access.
- Standard-output handling for failed commands
  (PR [#11](https://github.com/awslabs/tecli/pull/11)).
- Local test harness for the configure command
  (PR [#10](https://github.com/awslabs/tecli/pull/10)).
- Command-line argument inconsistencies vs. documented usage
  (PR [#9](https://github.com/awslabs/tecli/pull/9)).

## [0.4.0-alpha] - earlier

- Initial public release line under `awslabs/tecli`.
- Cobra-based CLI scaffolding, viper-backed credentials profile,
  embedded resources via the `box` package, and the `clencli` /
  `habits` integration that powers the README and Makefile.

[Unreleased]: https://github.com/awslabs/tecli/compare/v0.4.2-alpha...HEAD
[0.4.2-alpha]: https://github.com/awslabs/tecli/releases/tag/v0.4.2-alpha
[0.4.0-alpha]: https://github.com/awslabs/tecli/releases/tag/v0.4.0-alpha
