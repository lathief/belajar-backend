enum Status { Baru, Bayar }
struct Pesanan {
    kode: &'static str, // Kode contoh berupa literal yang hidup sepanjang program.
    status: Status,
}

fn main() {
    let daftar = [
        Pesanan { kode: "PS-1", status: Status::Baru },
        Pesanan { kode: "PS-2", status: Status::Bayar },
    ];
    for pesanan in daftar {
        let label = match pesanan.status {
            Status::Baru => "menunggu",
            Status::Bayar => "lunas",
        };
        println!("{} {label}", pesanan.kode);
    }
}
