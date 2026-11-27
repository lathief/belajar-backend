fn main() {
    let mut valid = 0;
    for jumlah in [0, 1, 4, 8] {
        // Kedua batas ikut diterima.
        if jumlah >= 1 && jumlah <= 5 {
            valid += 1;
        }
    }
    println!("valid={valid}");
}
