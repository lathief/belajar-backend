# Bab 07 — Koleksi dan pencarian

`Vec` dapat tumbuh dengan `push`. `iter()` meminjam tiap elemen. `HashMap` menghubungkan kunci dengan nilai. `get` menghasilkan `Option`, sehingga kunci yang tidak ada dapat ditangani secara eksplisit.

## Baca dan jalankan

Buka [main.rs](main.rs). Dari direktori `learn-rust/bab-07`, jalankan perintah berikut.

```sh
cargo run --quiet --bin bab-07
```

Keluaran program yang diharapkan

```text
jumlah item=3
stok=4
```

## Telusuri program

Setelah `push`, vektor berisi `[2, 3, 4]`. `len()` menghitung tiga elemen. Nilai stok 4 berasal dari pencarian kunci di map terpisah.

## Latihan mandiri

Cari `PS-9` dan ganti nilai default dengan cabang `match` yang mencetak `tidak ditemukan`.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-06/README.md) · [Bab berikutnya](../bab-08/README.md)
