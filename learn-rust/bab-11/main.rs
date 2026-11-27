#[derive(Clone, Copy)]
enum Status { Baru, Lunas, Dikirim, Batal }

impl Status {
    fn label(self) -> &'static str {
        match self {
            Self::Baru => "baru",
            Self::Lunas => "lunas",
            Self::Dikirim => "dikirim",
            Self::Batal => "batal",
        }
    }

    fn transisi(self, aksi: &str) -> Option<Self> {
        // Pasangkan keadaan awal dengan aksi sebelum memilih tujuan.
        match (self, aksi) {
            (Self::Baru, "bayar") => Some(Self::Lunas),
            (Self::Baru, "batal") => Some(Self::Batal),
            (Self::Lunas, "kirim") => Some(Self::Dikirim),
            _ => None,
        }
    }
}

fn main() {
    let mut status = Status::Baru;
    for aksi in ["kirim", "bayar", "batal", "kirim", "bayar"] {
        let tujuan = status.transisi(aksi);
        let berikut = tujuan.unwrap_or(status);
        println!("{} {aksi} -> {} ok={}", status.label(), berikut.label(), tujuan.is_some());
        status = berikut;
    }
}
