fn bersihkan(teks: &str) -> &str {
    // Potongan hasil masih mengacu pada buffer milik pemanggil.
    teks.trim()
}

fn main() {
    let kode = String::from(" PS-7 ");
    println!("kode={}", bersihkan(&kode));
    println!("panjang asli={}", kode.len());
}
