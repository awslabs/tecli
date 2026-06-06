# Contributing to TECLI

Thank you for your interest in contributing to TECLI. Whether it's a
bug report, a new command, a fix, or a docs improvement, we welcome
your help.

Please skim this document end-to-end before opening an issue or
pull request — it captures the toolchain, the branch flow, and the
test workflow that the maintainers expect.

## Reporting bugs / requesting features

Use the GitHub issue tracker. Before filing, please search existing
open and closed issues so we don't end up tracking the same problem
twice.

A good bug report typically contains:

- A reproducible test case or series of steps.
- The version of TECLI you're using (`tecli version`).
- The Go version, OS, and architecture you're running on.
- Any relevant config (`~/.tecli/credentials.yaml` profile shape, env
  vars), with secrets redacted.
- Anything unusual about your environment or deployment.

## Toolchain

TECLI now targets **Go 1.25** and uses
[`hashicorp/go-tfe`](https://github.com/hashicorp/go-tfe) `v1.108+`.

Set up a development environment:

```bash
# 1. Install Go 1.25+. https://go.dev/dl/
go version  # confirm 1.25+

# 2. Clone and build
git clone https://github.com/awslabs/tecli.git
cd tecli
go build ./...
```

There's also a [`.devcontainer/`](.devcontainer/) configuration if you
prefer VS Code's dev container experience.

The repo uses [pre-commit](https://pre-commit.com/) — install the
hooks once with:

```bash
pre-commit install
```

## Branch strategy

The default branch is `main`. Day-to-day work follows a simple
pattern:

| Branch               | Purpose                                                                                                |
| -------------------- | ------------------------------------------------------------------------------------------------------ |
| `main`               | Always green. Release-please cuts tags from this branch.                                               |
| Topic branches       | `feat/<short-name>`, `fix/<short-name>`, `docs/<short-name>` etc. Open a PR back to `main`.            |
| `chore/modernize-toolchain` | Long-running collector branch used during the June 2026 toolchain modernization (see PR #27). Sub-PRs target this branch first, then it merges to `main`. |

Use [Conventional Commits](https://www.conventionalcommits.org/) so
release-please can keep `CHANGELOG.md` accurate:

```
feat(workspace): add --auto-apply flag
fix(run): surface 422 error bodies instead of swallowing
docs: explain TFC_ORGANIZATION env var
chore(deps): bump go-tfe to v1.x
```

## How to run tests locally

The repository has two flavors of tests:

### Unit tests

```bash
go test ./...
```

This runs without external services and should always pass on a clean
checkout.

### Integration tests (Terraform Cloud)

The packages under `tests/commands/` exercise the live Terraform Cloud
API. They require credentials and will create real workspaces / runs in
your account.

```bash
# Either configure a profile (writes ~/.tecli/credentials.yaml)
tecli configure create

# ...or export environment variables
export TFC_ORGANIZATION=your-organization
export TFC_USER_TOKEN=your-user-token
export TFC_TEAM_TOKEN=your-team-token
export TFC_ORGANIZATION_TOKEN=your-organization-token

# Then run the integration suite via the Make target
make tecli/test
```

Integration tests are skipped automatically when the `TFC_*` env vars
and credentials profile are both absent.

> **Heads up:** The integration suite creates and tears down workspaces
> in the configured organization. Don't point it at a production org.

### Vet & format

```bash
go vet ./...
gofmt -s -w .
goimports -w .
```

`pre-commit` runs the same checks plus a few markdown linters.

## Sending us a pull request

1. Fork the repository.
2. Create a topic branch off the appropriate base (usually `main`,
   sometimes `chore/modernize-toolchain` during a modernization
   sweep).
3. Make focused changes — please don't reformat unrelated code in the
   same PR; it makes review harder.
4. Run `go build ./...`, `go vet ./...`, and any relevant tests.
5. Commit using Conventional Commits (see above).
6. Open the PR. Fill in the description: what's changing, why, how
   you tested it.
7. Stay engaged with CI failures and reviewer feedback.

GitHub has additional documentation on
[forking a repository](https://help.github.com/articles/fork-a-repo/)
and
[creating a pull request](https://help.github.com/articles/creating-a-pull-request/).

## Finding contributions to work on

The GitHub `help wanted` and `good first issue` labels are a good
starting point. The wiki under
<https://github.com/awslabs/tecli/wiki> also tracks longer-running
roadmap items.

## Code of Conduct

This project has adopted the
[Amazon Open Source Code of Conduct](https://aws.github.io/code-of-conduct).
For more information see the
[Code of Conduct FAQ](https://aws.github.io/code-of-conduct-faq) or
contact opensource-codeofconduct@amazon.com with any additional
questions or comments.

## Security issue notifications

If you discover a potential security issue in this project we ask that
you notify AWS/Amazon Security via our
[vulnerability reporting page](http://aws.amazon.com/security/vulnerability-reporting/).
**Please do not** create a public GitHub issue.

## Licensing

See the [LICENSE](LICENSE) file for our project's licensing. We may
ask you to confirm the licensing of your contribution. Larger
contributions may require a
[Contributor License Agreement (CLA)](http://en.wikipedia.org/wiki/Contributor_License_Agreement).
