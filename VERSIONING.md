## Branching strategy
This project has adopted [GitFlow](https://nvie.com/posts/a-successful-git-branching-model/) for its branching strategy model. 

## Release process
This project has adopted [Nebula Release Plugin]() for the release process.

### Explanation

`final` - Sets the version to the appropriate `<major>.<minor>.<patch>`, creates tag `v<major>.<minor>.<patch>`

`candidate` - Sets the version to the appropriate `<major>.<minor>.<patch>-rc.#`, creates tag `v<major>.<minor>.<patch>-rc.#` where `#` is the number of release candidates for this version produced so far. 1st 1.0.0 will be 1.0.0-rc.1, 2nd 1.0.0-rc.2 and so on.

`devSnapshot` - Sets the version to the appropriate `<major>.<minor>.<patch>-dev.#+<hash>`, does not create a tag. Where `#` is the number of commits since the last release and hash is the git hash of the current commit. If releasing a `devSnapshot` from a branch not listed in the releaseBranchPatterns and not excluded by excludeBranchPatterns the version will be `<major>.<minor>.<patch>-dev.#+<branchname>.<hash>`
You can use nebula.release.features.replaceDevWithImmutableSnapshot=true in your gradle.properties file to change pattern of version to <major>.<minor>.<patch>-snapshot.<timestamp>+<hash>. Where timestamp is UTC time in YYYYMMddHHmm format, ex. 201907052105 and hash is the git hash of the current commit. If releasing a immutableSnapshot from a branch not listed in the releaseBranchPatterns and not excluded by excludeBranchPatterns the version will be `<major>.<minor>.<patch>-snapshot.<timestamp>+<branchname>.<hash>`
