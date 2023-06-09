# Change log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0]

* Added `tables` module so that you can use it as a better abstraction to build up your entity changes.

  ```rust
  use substreams_database_change::tables::Tables;

  let mut tables = Tables::new();

  // Create a row (<entity_name>, <id>).[set(<column>, <value>), ...]
  tables
    .create_row("Factory", id)
    .set("poolCount", &bigint0)
    .set("txCount", &bigint0)
    .set("totalVolumeUSD", &bigdecimal0)
    .set("totalVolumeETH", &bigdecimal0)
    .set("totalFeesUSD", &bigdecimal0)
    .set("totalFeesETH", &bigdecimal0)
    .set("untrackedVolumeUSD", &bigdecimal0)
    .set("totalValueLockedUSD", &bigdecimal0)
    .set("totalValueLockedETH", &bigdecimal0)
    .set("totalValueLockedUSDUntracked", &bigdecimal0)
    .set("totalValueLockedETHUntracked", &bigdecimal0)
    .set("owner", &format!("0x{}", Hex(utils::ZERO_ADDRESS).to_string()));

  // Update a row (<entity_name>, <id>).[set(<column>, <value>), ...]
  tables
    .update_row("Bundle", "1").set("ethPriceUSD", &delta.new_value);
  ```

* Add composite primary keys for Protobuf with `oneof` and add `push_change_composite()` to `helpers.rs`.

* **Note** Changed database protobuf location path from ~`substreams/sink/database/v1`~ to `sf/substreams/sink/database/v1`, this is **not** a change of the Protobuf package id which is still `sf.substreams.sink.database.v1`. This will have an impact if you use `buf gen --exclude-path` as the path here is the Protobuf original path location, so add `sf/` if you were excluding `substreams/sink/database`.

* GitHub repository has been renamed to `substreams-sink-database-changes`, **note** the Rust crate's name is still `substreams-database-change` to avoid breaking change when importing the crate. We are planning a future rename of the crate, but that might come with a v2 of the data protocol.

## [1.0.1]

* Separate generic types for table name and private key

## [1.0.0]

* **BREAKING** Changed database proto package path from ~substreams.sink.database.v1~ to `sf.substreams.sink.database.v1`

## [0.2.0]

* **BREAKING** Updated `substreams` dependencies to `0.5.0`.

## [0.1.2]

* Added support for `substreams::Hex` type which converts to `string` in hexadecimal form.

## [0.1.1](https://github.com/streamingfast/substreams-sink-database-changes/releases/tag/v0.1.1)

* Added support for `prost::Timestamp` type.

* Made `AsString` public so you can implement on your own custom types.

## [0.1.0](https://github.com/streamingfast/substreams-sink-database-changes/releases/tag/v0.1.0)

* Added support for `u8`, `u16`, `u32`, `u64`, `i8`, `i16`, `i32`, `i64` types.

* Added possibility to record a `change` using `(new: AsString, old: AsString)` to simulate a value update.

* Added possibility to record a `change` using `(old: AsString, new: Option<Into<Typed>>)` to simulate a value deletion.

* Added possibility to record a `change` using `(old: Option<AsString>, new: AsString)` to simulate a value creation.

* Refactored to allow delta to be taken from any `AsString` which makes it much easier to extend when there is missing types.

* Introduced our own `AsString` type because `Into<String>` usage and `ToString` usage lead to Rust compiler errors on our desired change typings.

* Refactored to reduce amount of `clone` perform.

* Added `database.proto` containing proto message definitions
