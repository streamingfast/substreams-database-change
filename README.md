# Substreams Database Change

> Developer preview, used by database-like Substreams sinks like [substreams-sink-postgres](https://github.com/streamingfast/substreams-sink-postgres) and [substreams-sink-mongodb](https://github.com/streamingfast/substreams-sink-mongodb)

`substreams-database` crate contains all the definitions for the database changes which can be emitted by a substream.

### Install protoc

To be able to build proto files manually using the command line, you have to have protoc installed on your machine. Visit [here](https://grpc.io/docs/protoc-installation/) to install.

For Linux, using apt
```bash
apt install -y protobuf-compiler
protoc --version  # Ensure compiler version is 3+
```

For macOS, using Homebrew
```bash
brew install protobuf
protoc --version  # Ensure compiler version is 3+
```

Instead of having a `build.rs` file which will build the proto automatically on every `cargo build --release` you have to build the proto files manually.

Simply run `./gen/generate.sh`
