use std::collections::{HashSet, VecDeque};

struct Produsen {
    stok: u32,
    dipesan: bool,
    tertunda: bool,
}

impl Produsen {
    fn pesan(&mut self) {
        if !self.dipesan {
            // Satu penggantian state mewakili commit pesanan dan outbox.
            *self = Self { stok: self.stok - 1, dipesan: true, tertunda: true };
        }
    }
}

fn simulasi(respons_hilang: bool, mut relay_putus: bool, mut ack_hilang: bool) -> (u32, u32, u32) {
    let mut p = Produsen { stok: 5, dipesan: false, tertunda: false };
    p.pesan();
    if respons_hilang { p.pesan(); }

    let mut antrean = VecDeque::new();
    let mut publikasi = 0;
    while p.tertunda {
        antrean.push_back("A");
        publikasi += 1;
        if relay_putus {
            relay_putus = false;
            continue;
        }
        p.tertunda = false;
    }

    let mut inbox = HashSet::new();
    let mut paket = 0;
    while let Some(&id) = antrean.front() {
        // insert mengembalikan true tepat saat ID pertama kali tercatat.
        if inbox.insert(id) { paket += 1; }
        if ack_hilang {
            ack_hilang = false;
            continue;
        }
        antrean.pop_front();
    }
    (p.stok, publikasi, paket)
}

fn main() {
    let mut lulus = 0;
    for respons in [false, true] {
        for relay in [false, true] {
            for ack in [false, true] {
                let (stok, publikasi, paket) = simulasi(respons, relay, ack);
                assert_eq!((stok, publikasi, paket), (4, if relay { 2 } else { 1 }, 1));
                lulus += 1;
                if respons && relay && ack {
                    println!("stok={stok} publikasi={publikasi} paket={paket}");
                }
            }
        }
    }
    println!("skenario_lulus={lulus}");
}
