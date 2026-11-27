use std::collections::HashMap;

struct Bukti {
    jumlah: i32,
    hasil: &'static str,
}

struct Gudang {
    stok: i32,
    riwayat: HashMap<String, Bukti>,
}

impl Gudang {
    fn pesan(&mut self, id: &str, jumlah: i32) -> &'static str {
        if !(1..=5).contains(&jumlah) {
            return "jumlah_invalid";
        }

        // get meminjam bukti lama, sehingga isi map tetap utuh saat retry.
        if let Some(awal) = self.riwayat.get(id) {
            if awal.jumlah != jumlah {
                return "konflik";
            }
            return awal.hasil;
        }

        let hasil = if jumlah <= self.stok {
            self.stok -= jumlah;
            "diterima"
        } else {
            "stok_kurang"
        };
        self.riwayat.insert(id.to_owned(), Bukti { jumlah, hasil });
        hasil
    }
}

fn main() {
    let mut gudang = Gudang { stok: 5, riwayat: HashMap::new() };
    let urutan = [
        ("A", 3), ("B", 3), ("C", 0), ("C", 2),
        ("A", 3), ("A", 1), ("B", 3),
    ];
    for (id, jumlah) in urutan {
        let hasil = gudang.pesan(id, jumlah);
        println!("{id} {jumlah} {hasil} stok={}", gudang.stok);
    }
}
