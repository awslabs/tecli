# Releasing TECLI

This is the cookbook for cutting a new TECLI release. The repository
ships **two release-related GitHub Actions** plus a manual fallback:

| Workflow                                          | Trigger             | Purpose                                         |
| ------------------------------------------------- | ------------------- | ----------------------------------------------- |
| `.github/workflows/release.yml` (release-please)  | Push to `main`      | Opens/updates a release PR; tags on merge.      |
| `.github/workflows/publish.yml` (publish)         | Tag push (`v*`)     | Builds cross-platform binaries and a GH Release |
| `.github/workflows/build.yml`                     | Every push          | Smoke-build only; no release output.            |

## Standard release (release-please)

`release-please` watches Conventional Commits (`feat:`, `fix:`,
`chore:`...) on `main` and maintains a release PR titled something like
`chore(main): release tecli x.y.z`.

1. Land your work into `main` using Conventional Commits. The release
   PR is updated automatically.
2. Review the release PR — it bumps `box/resources/VERSION` and
   regenerates `CHANGELOG.md`. Edit the PR body if you want to override
   the auto-generated changelog.
3. Merge the release PR. release-please pushes a tag `vX.Y.Z` and
   creates a GitHub Release.
4. The tag push fires `publish.yml`, which compiles binaries for the
   matrix in [`Makefile`](../Makefile) (`tecli/compile` target) and
   uploads them to the GitHub Release as draft prereleases.
5. Confirm the binaries on the Releases page, untick **Draft** /
   **Pre-release** as appropriate, and publish.

### Conventional Commit cheatsheet

| Commit prefix     | Bumps          |
| ----------------- | -------------- |
| `feat: ...`       | minor          |
| `fix: ...`        | patch          |
| `feat!: ...` or `BREAKING CHANGE:` in trailer | major |
| `chore:`, `docs:`, `refactor:`, `test:`       | no bump        |

## Manual release fallback

If release-please is broken or you need an emergency tag:

```bash
# 1. Bump the embedded version
echo "0.5.0-alpha" > box/resources/VERSION
git add box/resources/VERSION
git commit -m "chore: release v0.5.0-alpha"

# 2. Update CHANGELOG.md by hand (Keep a Changelog format).
$EDITOR CHANGELOG.md
git add CHANGELOG.md
git commit -m "docs: changelog for v0.5.0-alpha"

# 3. Tag and push
git tag v0.5.0-alpha
git push origin main
git push origin v0.5.0-alpha
```

The tag push triggers `publish.yml`, which produces binaries and a
draft GitHub Release. From there:

```bash
# Optional: build binaries locally to double-check
make tecli/compile
ls dist/
```

Promote the GitHub Release from draft to published once the binaries
are verified.

## Versioning policy

- `0.x-alpha` while breaking changes are still expected.
- Promote to `0.x` once the public command surface is stable.
- Promote to `1.0.0` when SemVer guarantees can be made on the CLI's
  command names and flags.

## Pre-release checklist

- [ ] `go build ./...` passes locally on Go 1.25+.
- [ ] `go vet ./...` returns no new warnings.
- [ ] `make tecli/compile` succeeds (or rely on `publish.yml`).
- [ ] `CHANGELOG.md` lists the user-facing changes since the previous
      tag.
- [ ] `box/resources/VERSION` matches the new tag.
- [ ] `README.md`'s "Modernization" / install-version notes are
      consistent with the new toolchain.
