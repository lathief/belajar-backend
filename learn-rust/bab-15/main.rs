struct Ember {
    kapasitas: u32,
    token: u32,
    terakhir: u32,
}

impl Ember {
    fn ambil(&mut self, detik: u32) -> &'static str {
        // Selisih detik berasal dari jam monoton.
        self.token = self.token.saturating_add(detik - self.terakhir).min(self.kapasitas);
        self.terakhir = detik;
        if self.token == 0 { return "tunda"; }
        self.token -= 1;
        "lolos"
    }
}

fn main() {
    let mut e = Ember { kapasitas: 2, token: 2, terakhir: 0 };
    for t in [0, 0, 0, 1, 6, 6, 6] {
        let hasil = e.ambil(t);
        println!("t={t} {hasil} token={}", e.token);
    }
}
