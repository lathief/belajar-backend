fn kirim(respons: &[u16], deadline: u32) -> &'static str {
    let (mut waktu, mut jeda) = (0, 1);
    for (i, &kode) in respons.iter().enumerate() {
        println!("t={waktu} percobaan={} kode={kode}", i + 1);
        match kode {
            200 => return "sukses",
            503 => (),
            _ => return "permanen",
        }
        if i + 1 == respons.len() { return "habis"; }
        // Waktu tepat di deadline sudah di luar anggaran.
        if waktu + jeda >= deadline { return "deadline"; }
        waktu += jeda;
        jeda *= 2;
    }
    "habis"
}

fn main() {
    for (respons, batas) in [(&[503, 503, 200][..], 4), (&[503, 503, 200][..], 3), (&[400, 200][..], 4)] {
        let hasil = kirim(respons, batas);
        println!("hasil={hasil}");
    }
}
