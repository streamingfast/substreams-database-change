pub mod change;
pub mod helpers;
pub mod pb;
pub mod tables;

fn main() {
    let x = crate::pb::database::table_change::Operation::Unspecified;
    println!("{:?}", x);
}
