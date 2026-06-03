# Security Policy

## Reporting a Vulnerability

If you discover a potential security issue in this project we ask that you
**do not** create a public GitHub issue. Instead, please follow the AWS
vulnerability reporting process so the AWS Security team can triage and
coordinate a fix:

- Notify AWS Security via the
  [vulnerability reporting page](https://aws.amazon.com/security/vulnerability-reporting/).
- Optionally email **aws-security@amazon.com** directly. Please do **not**
  send sensitive details through GitHub issues, pull requests, or
  discussions.

When reporting, please include as much information as you can:

- A description of the issue and its impact.
- Steps to reproduce, including affected versions / commits if known.
- Any proof-of-concept code (please do not exploit production systems).
- Your name and contact info if you'd like credit in the release notes.

We will acknowledge receipt within a few business days and will keep you
informed of progress as we work toward a fix and disclosure.

## Supported Versions

This project follows the
[AWS Labs](https://github.com/awslabs) general support model: the latest
release on the default branch receives security fixes. Older tagged
releases are best-effort.

## Scope

`tecli` is a command-line client that makes authenticated calls to
[Terraform Cloud](https://app.terraform.io/) using a token supplied by the
operator. Reports about credential handling, command injection, supply-chain
issues (dependency tampering, build/release pipeline), or remote code
execution are particularly welcome.
