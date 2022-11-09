use crate::change::ToField;
use crate::pb::database::{DatabaseChanges, table_change::Operation, TableChange};

impl DatabaseChanges {
    pub fn push_change<V: AsRef<str>>(
        &mut self,
        table: V,
        pk: V,
        ordinal: u64,
        operation: Operation,
    ) -> &mut TableChange {
        let table_change = TableChange::new(table, pk, ordinal, operation);
        self.table_changes.push(table_change);
        return self.table_changes.last_mut().unwrap();
    }
}

impl TableChange {
    pub fn new<V: AsRef<str>>(
        entity: V,
        pk: V,
        ordinal: u64,
        operation: Operation,
    ) -> TableChange {
        TableChange {
            table: entity.as_ref().to_string(),
            pk: pk.as_ref().to_string(),
            ordinal,
            operation: operation as i32,
            fields: vec![],
        }
    }

    pub fn change<N: AsRef<str>, T: ToField>(&mut self, name: N, change: T) -> &mut TableChange {
        self.fields.push(change.to_field(name));
        self
    }
}
