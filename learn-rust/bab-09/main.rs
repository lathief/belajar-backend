use std::env;

fn main() {
    let mut unit: u64 = 0;
    let mut ditolak = 0;
    for teks in env::args().skip(1) {
        // Satu input buruk hanya menolak pesanan itu.
        match teks.parse::<u32>() {
            Ok(n) if (1..=5).contains(&n) => unit += u64::from(n),
            _ => ditolak += 1,
        }
    }
    println!("unit={unit} ditolak={ditolak}");
}
