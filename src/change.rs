use crate::pb::database::Field;
use substreams::pb::substreams::store_delta::Operation;
// use substreams::scalar::{BigDecimal, BigInt};
use substreams::scalar::{BigDecimal, BigInt};
use substreams::store::{DeltaBigDecimal, DeltaBigInt, DeltaBool, DeltaBytes, DeltaString};
use std::str;

pub trait ToField {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field;
}


// ---------- Uint32 ----------
impl ToField for u64 {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
        Field {
            name: name.as_ref().to_string(),
            new_value: self.to_string(),
            old_value: "".to_string(),
        }
    }
}


// ---------- Int32 ----------
impl ToField for i32 {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
        Field {
            name: name.as_ref().to_string(),
            new_value: self.to_string(),
            old_value: "".to_string(),
        }
    }
}

// impl ToField for  {
//     fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
//         let new_value = self.
//         let mut old_value: Option<Value> = None;
//
//         match Operation::from(self.operation) {
//             Operation::Update => {
//                 old_value = Some(Value {
//                     typed: Some(Typed::Int32(self.old_value)),
//                 });
//             }
//             Operation::Create => {}
//             _ => panic!("unsupported operation {:?}", self.operation),
//         }
//
//         Field {
//             name: name.as_ref().to_string(),
//             new_value,
//             old_value,
//         }
//     }
// }

// ---------- BigDecimal ----------
impl ToField for BigDecimal {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
        Field {
            name: name.as_ref().to_string(),
            new_value: self.to_string(),
            old_value: "".to_string(),
        }
    }
}

impl ToField for DeltaBigDecimal {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
        let new_value = self.new_value.to_string();
        let mut old_value = "".to_string();

        match Operation::from(self.operation) {
            Operation::Update => {
                old_value = self.old_value.to_string().clone()
            }
            Operation::Create => {}
            _ => panic!("unsupported operation {:?}", self.operation),
        }

        Field {
            name: name.as_ref().to_string(),
            new_value,
            old_value,
        }
    }
}

// ---------- BigInt ----------
impl ToField for BigInt {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
        Field {
            name: name.as_ref().to_string(),
            new_value: self.to_string(),
            old_value: "".to_string(),
        }
    }
}

impl ToField for DeltaBigInt {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
        let new_value = self.new_value.to_string();
        let mut old_value = "".to_string();

        match Operation::from(self.operation) {
            Operation::Update => {
                old_value = self.old_value.to_string()
            }
            Operation::Create => {}
            _ => panic!("unsupported operation {:?}", self.operation),
        }

        Field {
            name: name.as_ref().to_string(),
            new_value,
            old_value,
        }
    }
}

// ---------- String ----------
impl ToField for String {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
        Field {
            name: name.as_ref().to_string(),
            new_value: self.to_string(),
            old_value: "".to_string(),
        }
    }
}

impl ToField for DeltaString {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
        let new_value = self.new_value.to_string();
        let mut old_value = "".to_string();

        match Operation::from(self.operation) {
            Operation::Update => {
                old_value = self.old_value.to_string();
            }
            Operation::Create => {}
            _ => panic!("unsupported operation {:?}", self.operation),
        }

        Field {
            name: name.as_ref().to_string(),
            new_value,
            old_value,
        }
    }
}

// ---------- Bytes ----------
impl ToField for Vec<u8> {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
        Field {
            name: name.as_ref().to_string(),
            new_value: base64::encode(self.clone()),
            old_value: "".to_string(),
        }
    }
}

impl ToField for DeltaBytes {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
        let new_value = base64::encode(self.new_value.clone());
        let mut old_value = "".to_string();

        match Operation::from(self.operation) {
            Operation::Update => {
                old_value = base64::encode(self.old_value.clone())
            }
            Operation::Create => {}
            _ => panic!("unsupported operation {:?}", self.operation),
        }

        Field {
            name: name.as_ref().to_string(),
            new_value,
            old_value,
        }
    }
}

// ---------- BoolChange ----------
impl ToField for bool {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
        Field {
            name: name.as_ref().to_string(),
            new_value: self.to_string(),
            old_value: "".to_string(),
        }
    }
}

impl ToField for DeltaBool {
    fn to_field<N: AsRef<str>>(&self, name: N) -> Field {
        let new_value = self.new_value.to_string();
        let mut old_value = "".to_string();

        match Operation::from(self.operation) {
            Operation::Update => {
                old_value = self.old_value.to_string();
            }
            Operation::Create => {}
            _ => panic!("unsupported operation {:?}", self.operation),
        }

        Field {
            name: name.as_ref().to_string(),
            new_value,
            old_value,
        }
    }
}


#[cfg(test)]
mod test {
    use substreams::pb::substreams::store_delta::Operation;
    use crate::pb::database::Field;
    use crate::change::ToField;
    use substreams::scalar::{BigDecimal, BigInt};
    use substreams::store::{DeltaBigDecimal, DeltaBigInt, DeltaBool, DeltaBytes, DeltaString};

    const FIELD_NAME: &str = "field.name.1";

    #[test]
    fn i32_change() {
        let i32_change: i32 = 1;
        assert_eq!(
            create_expected_field(FIELD_NAME, None, Some("1".to_string())),
            i32_change.to_field(FIELD_NAME)
        );
    }

    #[test]
    fn big_decimal_change() {
        let bd_change = BigDecimal::from(1 as i32);
        assert_eq!(
            create_expected_field(FIELD_NAME, None, Some("1".to_string())),
            bd_change.to_field(FIELD_NAME)
        );
    }

    #[test]
    fn delta_big_decimal_change() {
        let delta = DeltaBigDecimal {
            operation: Operation::Update,
            ordinal: 0,
            key: "change".to_string(),
            old_value: BigDecimal::from(10),
            new_value: BigDecimal::from(20),
        };

        assert_eq!(
            create_expected_field(
                FIELD_NAME,
                Some("10".to_string()),
                Some("20".to_string())
            ),
            delta.to_field(FIELD_NAME)
        );
    }

    #[test]
    fn big_int_change() {
        let bi_change = BigInt::from(1 as i32);
        assert_eq!(
            create_expected_field(FIELD_NAME, None, Some("1".to_string())),
            bi_change.to_field(FIELD_NAME)
        );
    }

    #[test]
    fn delta_big_int_change() {
        let delta = DeltaBigInt {
            operation: Operation::Update,
            ordinal: 0,
            key: "change".to_string(),
            old_value: BigInt::from(10),
            new_value: BigInt::from(20),
        };

        assert_eq!(
            create_expected_field(
                FIELD_NAME,
                Some("10".to_string()),
                Some("20".to_string())
            ),
            delta.to_field(FIELD_NAME)
        );
    }

    #[test]
    fn string_change() {
        let string_change = String::from("string");
        assert_eq!(
            create_expected_field(FIELD_NAME, None, Some("string".to_string())),
            string_change.to_field(FIELD_NAME)
        );
    }

    #[test]
    fn delta_string_change() {
        let delta = DeltaString {
            operation: Operation::Update,
            ordinal: 0,
            key: "change".to_string(),
            old_value: String::from("string1"),
            new_value: String::from("string2"),
        };

        assert_eq!(
            create_expected_field(
                FIELD_NAME,
                Some("string1".to_string()),
                Some("string2".to_string())
            ),
            delta.to_field(FIELD_NAME)
        );
    }

    #[test]
    fn bytes_change() {
        let bytes_change = Vec::from("bytes");
        assert_eq!(
            create_expected_field(FIELD_NAME, None, Some("Ynl0ZXM=".to_string())),
            bytes_change.to_field(FIELD_NAME)
        );
    }

    #[test]
    fn delta_bytes_change() {
        let delta = DeltaBytes {
            operation: Operation::Update,
            ordinal: 0,
            key: "change".to_string(),
            old_value: Vec::from("bytes1"),
            new_value: Vec::from("bytes2"),
        };

        assert_eq!(
            create_expected_field(
                FIELD_NAME,
                Some("Ynl0ZXMx".to_string()),
                Some("Ynl0ZXMy".to_string())
            ),
            delta.to_field(FIELD_NAME)
        );
    }

    #[test]
    fn bool_change() {
        let bool_change = true;
        assert_eq!(
            create_expected_field(FIELD_NAME, None, Some("true".to_string())),
            bool_change.to_field(FIELD_NAME)
        );
    }

    #[test]
    fn delta_bool_change() {
        let delta = DeltaBool {
            operation: Operation::Update,
            ordinal: 0,
            key: "change".to_string(),
            old_value: true,
            new_value: false,
        };

        assert_eq!(
            create_expected_field(
                FIELD_NAME,
                Some(true.to_string()),
                Some(false.to_string()),
            ),
            delta.to_field(FIELD_NAME)
        );
    }

    fn create_expected_field<N: AsRef<str>>(
        name: N,
        old_value: Option<String>,
        new_value: Option<String>,
    ) -> Field {
        let mut field = Field {
            name: name.as_ref().to_string(),
            ..Default::default()
        };
        if old_value.is_some() {
            field.old_value = old_value.unwrap()
        }
        if new_value.is_some() {
            field.new_value = new_value.unwrap()
        }
        field
    }
}
