struct Alamat {
    kota: String,
    versi: u32,
}

impl Alamat {
    fn ubah(&mut self, dibaca: u32, kota: &str) -> &'static str {
        if self.versi != dibaca {
            return "konflik";
        }
        // Simpan teks milik objek hanya setelah versi cocok.
        self.kota = kota.to_owned();
        self.versi += 1;
        "tersimpan"
    }
}

fn lapor(a: &mut Alamat, pelaku: &str, versi: u32, kota: &str) {
    let hasil = a.ubah(versi, kota);
    println!("{pelaku} {hasil} kota={} versi={}", a.kota, a.versi);
}

fn main() {
    let mut a = Alamat { kota: "Bandung".into(), versi: 1 };
    let (baca_a, baca_b) = (a.versi, a.versi);
    lapor(&mut a, "A", baca_a, "Solo");
    lapor(&mut a, "B", baca_b, "Bogor");
    let setelah_konfirmasi = a.versi;
    lapor(&mut a, "B", setelah_konfirmasi, "Bogor");
}
