use std::collections::HashMap;

fn main() {
    let mut item = vec![2, 3];
    item.push(4);
    let stok = HashMap::from([("PS-1", 4)]);
    // copied mengubah Option<&i32> menjadi Option<i32>.
    let tersedia = stok.get("PS-1").copied().unwrap_or(0);
    println!("jumlah item={}", item.len());
    println!("stok={tersedia}");
}
