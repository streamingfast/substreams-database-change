use crate::pb::database::{table_change::Operation, DatabaseChanges, Field, TableChange};
use std::collections::{BTreeMap, HashMap};
use substreams::{
    scalar::{BigDecimal, BigInt},
    Hex,
};

#[derive(Debug)]
pub struct Tables {
    // Map from table name to the primary keys within that table
    pub tables: HashMap<String, Rows>,
}

impl Tables {
    pub fn new() -> Self {
        Tables {
            tables: HashMap::new(),
        }
    }

    pub fn create_row<K: Into<PrimaryKey>>(&mut self, table: &str, key: K) -> &mut Row {
        let rows = self.tables.entry(table.to_string()).or_insert(Rows::new());
        let k = key.into();
        let key_debug = format!("{:?}", k);
        let row = rows.pks.entry(k).or_insert(Row::new());
        match row.operation {
            Operation::Unspecified => {
                row.operation = Operation::Create;
            }
            Operation::Create => {}
            Operation::Update => {
                panic!("cannot create a row that was marked for update")
            }
            Operation::Delete => {
                panic!(
                    "cannot create a row after a scheduled delete operation - table: {} key: {}",
                    table, key_debug,
                )
            }
        }
        row
    }

    pub fn update_row<K: Into<PrimaryKey>>(&mut self, table: &str, key: K) -> &mut Row {
        let rows = self.tables.entry(table.to_string()).or_insert(Rows::new());
        let k = key.into();
        let key_debug = format!("{:?}", k);
        let row = rows.pks.entry(k).or_insert(Row::new());
        match row.operation {
            Operation::Unspecified => {
                row.operation = Operation::Update;
            }
            Operation::Create => {}
            Operation::Update => {}
            Operation::Delete => {
                panic!(
                    "cannot update a row after a scheduled delete operation - table: {} key: {}",
                    table, key_debug,
                )
            }
        }
        row
    }

    pub fn delete_row<K: Into<PrimaryKey>>(&mut self, table: &str, key: PrimaryKey) -> &mut Row {
        let rows = self.tables.entry(table.to_string()).or_insert(Rows::new());
        let row = rows.pks.entry(key.into()).or_insert(Row::new());
        match row.operation {
            Operation::Unspecified => {
                row.operation = Operation::Delete;
            }
            Operation::Create => {
                // simply clear the thing
                row.operation = Operation::Unspecified;
                row.columns = HashMap::new();
            }
            Operation::Update => {
                row.columns = HashMap::new();
            }
            Operation::Delete => {}
        }
        row.operation = Operation::Delete;
        row.columns = HashMap::new();
        row
    }

    // Convert Tables into an DatabaseChanges protobuf object
    pub fn to_database_changes(self) -> DatabaseChanges {
        let mut changes = DatabaseChanges::default();

        for (table, rows) in self.tables.into_iter() {
            for (pk, row) in rows.pks.into_iter() {
                if row.operation == Operation::Unspecified {
                    continue;
                }

                let mut change = match pk {
                    PrimaryKey::Single(pk) => TableChange::new(table.clone(), pk, 0, row.operation),
                    PrimaryKey::Composite(keys) => TableChange::new_composite(
                        table.clone(),
                        keys.into_iter().collect(),
                        0,
                        row.operation,
                    ),
                };

                for (field, value) in row.columns.into_iter() {
                    change.fields.push(Field {
                        name: field,
                        new_value: value,
                        old_value: "".to_string(),
                    });
                }

                changes.table_changes.push(change);
            }
        }

        changes
    }
}

#[derive(Hash, Debug, Eq, PartialEq)]
pub enum PrimaryKey {
    Single(String),
    Composite(BTreeMap<String, String>),
}

impl From<&str> for PrimaryKey {
    fn from(x: &str) -> Self {
        Self::Single(x.to_string())
    }
}

impl From<&String> for PrimaryKey {
    fn from(x: &String) -> Self {
        Self::Single(x.clone())
    }
}

impl From<String> for PrimaryKey {
    fn from(x: String) -> Self {
        Self::Single(x)
    }
}

impl<K: AsRef<str>, const N: usize> From<[(K, String); N]> for PrimaryKey {
    fn from(arr: [(K, String); N]) -> Self {
        if N == 0 {
            return Self::Composite(BTreeMap::new());
        }

        let string_arr = arr.map(|(k, v)| (k.as_ref().to_string(), v));
        Self::Composite(BTreeMap::from(string_arr))
    }
}

#[derive(Debug)]
pub struct Rows {
    // Map of primary keys within this table, to the fields within
    pks: HashMap<PrimaryKey, Row>,
}

impl Rows {
    pub fn new() -> Self {
        Rows {
            pks: HashMap::new(),
        }
    }
}

#[derive(Debug)]
pub struct Row {
    // Verify that we don't try to delete the same row as we're creating it
    pub operation: Operation,
    // Map of field name to its last change
    pub columns: HashMap<String, String>,
    // Finalized: Last update or delete
    pub finalized: bool,
}

impl Row {
    pub fn new() -> Self {
        Row {
            operation: Operation::Unspecified,
            columns: HashMap::new(),
            finalized: false,
        }
    }

    pub fn set<T: ToDatabaseValue>(&mut self, name: &str, value: T) -> &mut Self {
        if self.operation == Operation::Delete {
            panic!("cannot set fields on a delete operation")
        }
        self.columns.insert(name.to_string(), value.to_value());
        self
    }

    pub fn set_raw(&mut self, name: &str, value: String) -> &mut Self {
        self.columns.insert(name.to_string(), value);
        self
    }

    /// Set a field to an array of values compatible with PostgresSQL database,
    /// this method is currently experimental and hidden as we plan to support
    /// array natively in the model.
    ///
    /// For now, this method should be used with great care as it ties the model
    /// to the database implementation.
    #[doc(hidden)]
    pub fn set_psql_array<T: ToDatabaseValue>(&mut self, name: &str, value: Vec<T>) -> &mut Row {
        if self.operation == Operation::Delete {
            panic!("cannot set fields on a delete operation")
        }

        let values = value
            .into_iter()
            .map(|x| x.to_value())
            .collect::<Vec<_>>()
            .join(",");

        self.columns.insert(name.to_string(), format!("'{{{}}}'", values));
        self
    }

    /// Set a field to an array of values compatible with Clickhouse database,
    /// this method is currently experimental and hidden as we plan to support
    /// array natively in the model.
    ///
    /// For now, this method should be used with great care as it ties the model
    /// to the database implementation.
    #[doc(hidden)]
    pub fn set_clickhouse_array<T: ToDatabaseValue>(&mut self, name: &str, value: Vec<T>) -> &mut Row {
        if self.operation == Operation::Delete {
            panic!("cannot set fields on a delete operation")
        }

        let values = value
            .into_iter()
            .map(|x| x.to_value())
            .collect::<Vec<_>>()
            .join(",");

        self.columns.insert(name.to_string(), format!("[{}]", values));
        self
    }
}

macro_rules! impl_to_database_value_proxy_to_ref {
    ($name:ty) => {
        impl ToDatabaseValue for $name {
            fn to_value(self) -> String {
                ToDatabaseValue::to_value(&self)
            }
        }
    };
}

macro_rules! impl_to_database_value_proxy_to_string {
    ($name:ty) => {
        impl ToDatabaseValue for $name {
            fn to_value(self) -> String {
                ToString::to_string(&self)
            }
        }
    };
}

pub trait ToDatabaseValue {
    fn to_value(self) -> String;
}

impl_to_database_value_proxy_to_string!(i8);
impl_to_database_value_proxy_to_string!(i16);
impl_to_database_value_proxy_to_string!(i32);
impl_to_database_value_proxy_to_string!(i64);
impl_to_database_value_proxy_to_string!(u8);
impl_to_database_value_proxy_to_string!(u16);
impl_to_database_value_proxy_to_string!(u32);
impl_to_database_value_proxy_to_string!(u64);
impl_to_database_value_proxy_to_string!(bool);
impl_to_database_value_proxy_to_string!(::prost_types::Timestamp);
impl_to_database_value_proxy_to_string!(&::prost_types::Timestamp);
impl_to_database_value_proxy_to_string!(&str);
impl_to_database_value_proxy_to_string!(BigDecimal);
impl_to_database_value_proxy_to_string!(&BigDecimal);
impl_to_database_value_proxy_to_string!(BigInt);
impl_to_database_value_proxy_to_string!(&BigInt);

impl_to_database_value_proxy_to_ref!(Vec<u8>);

impl ToDatabaseValue for &String {
    fn to_value(self) -> String {
        self.clone()
    }
}

impl ToDatabaseValue for String {
    fn to_value(self) -> String {
        self
    }
}

impl ToDatabaseValue for &Vec<u8> {
    fn to_value(self) -> String {
        Hex::encode(self)
    }
}

impl<T: AsRef<[u8]>> ToDatabaseValue for Hex<T> {
    fn to_value(self) -> String {
        ToString::to_string(&self)
    }
}

impl<T: AsRef<[u8]>> ToDatabaseValue for &Hex<T> {
    fn to_value(self) -> String {
        ToString::to_string(self)
    }
}

#[cfg(test)]
mod test {
    use crate::pb::database::table_change::PrimaryKey;
    use crate::pb::database::CompositePrimaryKey;
    use crate::pb::database::{DatabaseChanges, TableChange};
    use crate::tables::PrimaryKey as TablesPrimaryKey;
    use crate::tables::Tables;
    use crate::tables::ToDatabaseValue;
    use std::collections::HashMap;

    #[test]
    fn to_database_value_proto_timestamp() {
        assert_eq!(
            ToDatabaseValue::to_value(::prost_types::Timestamp {
                seconds: 60 * 60 + 60 + 1,
                nanos: 1
            }),
            "1970-01-01T01:01:01.000000001Z"
        );
    }

    #[test]
    fn create_row_single_pk_direct() {
        let mut tables = Tables::new();
        tables.create_row("myevent", TablesPrimaryKey::Single("myhash".to_string()));

        assert_eq!(
            tables.to_database_changes(),
            DatabaseChanges {
                table_changes: [TableChange {
                    table: "myevent".to_string(),
                    ordinal: 0,
                    operation: 1,
                    fields: [].into(),
                    primary_key: Some(PrimaryKey::Pk("myhash".to_string())),
                }]
                .to_vec(),
            }
        );
    }

    #[test]
    fn create_row_single_pk() {
        let mut tables = Tables::new();
        tables.create_row("myevent", "myhash");

        assert_eq!(
            tables.to_database_changes(),
            DatabaseChanges {
                table_changes: [TableChange {
                    table: "myevent".to_string(),
                    ordinal: 0,
                    operation: 1,
                    fields: [].into(),
                    primary_key: Some(PrimaryKey::Pk("myhash".to_string())),
                }]
                .to_vec(),
            }
        );
    }

    #[test]
    fn create_row_composite_pk() {
        let mut tables = Tables::new();
        tables.create_row(
            "myevent",
            [
                ("evt_tx_hash", "hello".to_string()),
                ("evt_index", "world".to_string()),
            ],
        );

        assert_eq!(
            tables.to_database_changes(),
            DatabaseChanges {
                table_changes: [TableChange {
                    table: "myevent".to_string(),
                    ordinal: 0,
                    operation: 1,
                    fields: [].into(),
                    primary_key: Some(PrimaryKey::CompositePk(CompositePrimaryKey {
                        keys: HashMap::from([
                            ("evt_tx_hash".to_string(), "hello".to_string()),
                            ("evt_index".to_string(), "world".to_string())
                        ])
                    }))
                }]
                .to_vec(),
            }
        );
    }
}
