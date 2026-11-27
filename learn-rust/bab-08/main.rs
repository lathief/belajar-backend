fn jumlah(teks: &str) -> Result<i32, &'static str> {
    let nilai: i32 = teks.parse().map_err(|_| "bukan angka")?;
    // Parsing sukses juga dapat menghasilkan angka nol.
    if nilai <= 0 { return Err("harus positif"); }
    Ok(nilai)
}

fn main() {
    for teks in ["2", "0", "x"] {
        match jumlah(teks) {
            Ok(n) => println!("jumlah={n}"),
            Err(alasan) => println!("{alasan}"),
        }
    }
}
