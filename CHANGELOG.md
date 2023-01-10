# Change log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0]

* **BREAKING** Changed database proto package path from ~substreams.sink.database.v1~ to `sf.substreams.sink.database.v1`

## [0.2.0]

* **BREAKING** Updated `substreams` dependencies to `0.5.0`.

## [0.1.2]

* Added support for `substreams::Hex` type which converts to `string` in hexadecimal form.

## [0.1.1](https://github.com/streamingfast/substreams-database-change/releases/tag/v0.1.1)

* Added support for `prost::Timestamp` type.

* Made `AsString` public so you can implement on your own custom types.

## [0.1.0](https://github.com/streamingfast/substreams-database-change/releases/tag/v0.1.0)

* Added support for `u8`, `u16`, `u32`, `u64`, `i8`, `i16`, `i32`, `i64` types.

* Added possibility to record a `change` using `(new: AsString, old: AsString)` to simulate a value update.

* Added possibility to record a `change` using `(old: AsString, new: Option<Into<Typed>>)` to simulate a value deletion.

* Added possibility to record a `change` using `(old: Option<AsString>, new: AsString)` to simulate a value creation.

* Refactored to allow delta to be taken from any `AsString` which makes it much easier to extend when there is missing types.

* Introduced our own `AsString` type because `Into<String>` usage and `ToString` usage lead to Rust compiler errors on our desired change typings.

* Refactored to reduce amount of `clone` perform.

* Added `database.proto` containing proto message definitions
