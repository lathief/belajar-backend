#[derive(Default)]
struct Tugas {
    generasi: u64,
    kedaluwarsa: u64,
    selesai: bool,
}

impl Tugas {
    fn klaim(&mut self, kini: u64) -> Option<u64> {
        if self.selesai || kini < self.kedaluwarsa { return None; }
        self.generasi += 1;
        self.kedaluwarsa = kini + 5;
        Some(self.generasi)
    }

    fn simpan(&mut self, token: u64, kini: u64) -> &'static str {
        // Finalisasi memeriksa token, keadaan terminal, dan waktu.
        if token == 0 || self.selesai || token != self.generasi || kini >= self.kedaluwarsa {
            return "ditolak";
        }
        self.selesai = true;
        "tersimpan"
    }
}

fn main() {
    let mut tugas = Tugas::default();
    let a = tugas.klaim(0).expect("tugas tersedia");
    println!("klaim_dini={}", tugas.klaim(4).is_some());
    println!("batas={}", tugas.simpan(a, 5));
    let b = tugas.klaim(5).expect("lease lama sudah berakhir");
    println!("lama={}", tugas.simpan(a, 6));
    println!("baru={}", tugas.simpan(b, 6));
    println!("ulang={}", tugas.simpan(b, 7));
}
