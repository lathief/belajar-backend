fn subtotal(harga: u32, jumlah: u32) -> u32 {
    // Ekspresi ini menjadi hasil fungsi.
    harga * jumlah
}

fn main() {
    println!("total={}", subtotal(5000, 3));
}
