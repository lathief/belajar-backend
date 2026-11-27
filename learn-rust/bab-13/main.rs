#[derive(Default)]
struct Ringkasan {
    posisi: u32,
    stok: i32,
}

impl Ringkasan {
    fn terapkan(&mut self, nomor: u32, delta: i32) -> &'static str {
        if nomor <= self.posisi { return "duplikat"; }
        if nomor != self.posisi + 1 { return "celah"; }
        // Perbarui posisi dan stok pada cabang sukses yang sama.
        self.stok += delta;
        self.posisi = nomor;
        "diterapkan"
    }
}

fn main() {
    let mut r = Ringkasan::default();
    for (nomor, delta) in [(1, 5), (3, -2), (2, -1), (2, -1), (3, -2)] {
        let hasil = r.terapkan(nomor, delta);
        println!("{nomor} {hasil} posisi={} stok={}", r.posisi, r.stok);
    }
}
