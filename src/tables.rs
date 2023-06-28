use crate::pb::database::{table_change::Operation, DatabaseChanges, Field, TableChange};
use std::collections::HashMap;
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

    pub fn create_row<K: AsRef<str>>(&mut self, table: &str, key: K) -> &mut Row {
        let rows = self.tables.entry(table.to_string()).or_insert(Rows::new());
        let row = rows
            .pks
            .entry(key.as_ref().to_string())
            .or_insert(Row::new());
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
                    table,
                    key.as_ref().to_string()
                )
            }
        }
        row
    }

    pub fn update_row<K: AsRef<str>>(&mut self, table: &str, key: K) -> &mut Row {
        let rows = self.tables.entry(table.to_string()).or_insert(Rows::new());
        let row = rows
            .pks
            .entry(key.as_ref().to_string())
            .or_insert(Row::new());
        match row.operation {
            Operation::Unspecified => {
                row.operation = Operation::Update;
            }
            Operation::Create => {}
            Operation::Update => {}
            Operation::Delete => {
                panic!(
                    "cannot update a row after a scheduled delete operation - table: {} key: {}",
                    table,
                    key.as_ref().to_string()
                )
            }
        }
        row
    }

    pub fn delete_row<K: AsRef<str>>(&mut self, table: &str, key: K) -> &mut Row {
        let rows = self.tables.entry(table.to_string()).or_insert(Rows::new());
        let row = rows
            .pks
            .entry(key.as_ref().to_string())
            .or_insert(Row::new());
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

                let mut change = TableChange::new(table.clone(), pk, 0, row.operation);
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

#[derive(Debug)]
pub struct Rows {
    // Map of primary keys within this table, to the fields within
    pub pks: HashMap<String, Row>,
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
    use crate::tables::ToDatabaseValue;

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
}
