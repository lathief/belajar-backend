fn main() {
    let harga: u32 = 5000;
    let mut jumlah = 2;
    // Barang tambahan masuk sebelum total dihitung.
    jumlah += 1;
    let subtotal = harga * jumlah;
    println!("subtotal={subtotal}");
}
