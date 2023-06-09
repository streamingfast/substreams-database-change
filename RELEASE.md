## Release Process

### Instructions

> **Warning** Do not forget to replace `${version}` by your real version like `0.1.3` in the commands below!  (`export version=0.1.3`)

You will need [sfreleaser](https://github.com/streamingfast/sfreleaser) (install from source using Golang with `go install github.com/streamingfast/sfreleaser/cmd/sfrelease@latest`) to perform the release process.

- Find & replace all occurrences of Regex `^version = "[^"]+"` in all `Cargo.toml` files to `version = "${version}"` and update the [CHANGELOG.md](CHANGELOG.md) to update the `## Unreleased` header to become `## ${version}`. Also don't forget to update `substreams.yaml` version.

  Using [sd](https://github.com/chmln/sd):

  ```bash
  find . -type f -name Cargo.toml -not -path "./target/*" | xargs -n1 sd '^version = "[^"]+"' "version = \"${version}\""
  sd '## Unreleased' "## v{version}" CHANGELOG.md
  sd 'version: v.*' "version: v${version}" substreams.yaml
  ```

- Perform a `cargo check` so that `Cargo.lock` is properly updated and run the tests `cargo test --target <your_machine_targe>`.

- Commit everything with message `Preparing release of ${version}`.

  ```bash
  git add -A . && git commit -m "Preparing release of ${version}"
  ```

- Perform release

  ```bash
  sfreleaser release
  ```

#### Easy script

```bash
# Update to correct version!
export version=0.1.3

find . -type f -name Cargo.toml -not -path "./target/*" | xargs -n1 sd '^version = "[^"]+"' "version = \"${version}\""
sd '## Unreleased' "## v{version}" CHANGELOG.md
sd 'version: v.*' "version: v${version}" substreams.yaml

cargo check

git add -A . && git commit -m "Preparing release of ${version}"

sfreleaser release v${version}
```