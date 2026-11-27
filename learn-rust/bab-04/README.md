# Bab 04 — Fungsi dan nilai balik

Parameter ditulis sebagai `nama: Tipe`. Panah `->` menyatakan tipe hasil. Ekspresi terakhir tanpa titik koma menjadi nilai balik. Pisahkan perhitungan dari pencetakan agar aturan harga bisa dipakai berulang.

## Baca dan jalankan

Buka [main.rs](main.rs). Dari direktori `learn-rust/bab-04`, jalankan perintah berikut.

```sh
cargo run --quiet --bin bab-04
```

Keluaran program yang diharapkan

```text
total=15000
```

## Telusuri program

Saat `subtotal(5000, 3)` dipanggil, `harga` bernilai 5000 dan `jumlah` bernilai 3. Menambahkan `;` setelah perkalian membuat blok menghasilkan `()`, sehingga tipe hasilnya tidak cocok dengan `u32`.

## Latihan mandiri

Tambahkan parameter biaya tetap dan kirim 1000. Untuk input semula hasilnya harus 16000.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-03/README.md) · [Bab berikutnya](../bab-05/README.md)
