# Substreams Database Change

[<img alt="github" src="https://img.shields.io/badge/Github-substreams.database-8da0cb?style=for-the-badge&logo=github" height="20">](https://github.com/streamingfast/substreams-sink-database-changes)
[<img alt="crates.io" src="https://img.shields.io/crates/v/substreams-sink-database-changes.svg?style=for-the-badge&color=fc8d62&logo=rust" height="20">](https://crates.io/crates/substreams-sink-database-changes)
[<img alt="docs.rs" src="https://img.shields.io/badge/docs.rs-substreams.database-66c2a5?style=for-the-badge&labelColor=555555&logo=docs.rs" height="20">](https://docs.rs/substreams-sink-database-changes)
[<img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/streamingfast/substreams-sink-database-changes/ci.yml?branch=develop&style=for-the-badge" height="20">](https://github.com/streamingfast/substreams-sink-database-changes/actions?query=branch%3Adevelop)

> `substreams-sink-database-changes` contains all the definitions for database changes which can be emitted by a substream.

## Used by

- [substreams-sink-postgres](https://github.com/streamingfast/substreams-sink-postgres)
- [substreams-sink-mongodb](https://github.com/streamingfast/substreams-sink-mongodb)

## Install    

```bash
$ cargo add substreams-sink-database-changes
```

## Quickstart

**Cargo.toml**

```toml
[dependencies]
substreams = "0.5"
substreams-sink-database-changes = "1.0"
```

**src/lib.rs**

```rust
use substreams::errors::Error;
use substreams_sink_database_changes::pb::database::{DatabaseChanges, table_change::Operation};

#[substreams::handlers::map]
fn db_out(
    ... some stores ...
) -> Result<DatabaseChanges, Error> {
    // Initialize Database Changes container
    let mut database_changes: DatabaseChanges = Default::default();

    // Push change
    database_changes.push_change("transfer", "primary-key", 0, Operation::Create)
        .change("key1", ("previous1", "value1"))
        .change("key2", ("previous2", "value2"))

    Ok(database_changes)
}
```
