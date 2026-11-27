# Bab 09 — Proyek mini rekapan pesanan

Program CLI menerima satu jumlah pesanan per argumen. `skip(1)` melewati nama program. `match` dengan guard `if` memisahkan angka dalam rentang 1–5 dari input lain. Rekap menyimpan jumlah unit diterima dan jumlah entri ditolak.

## Baca dan jalankan

Buka [main.rs](main.rs). Dari direktori `learn-rust/bab-09`, jalankan perintah berikut.

```sh
cargo run --quiet --bin bab-09 -- 2 0 x 3
```

Keluaran program yang diharapkan

```text
unit=5 ditolak=2
```

## Telusuri program

Argumen `2` dan `3` masuk cabang `Ok(n)` yang memenuhi guard. `0` gagal pada guard dan `x` gagal saat parsing. Keduanya masuk cabang `_`.

## Latihan mandiri

Tambahkan hitungan pesanan valid. Untuk `2 0 x 3`, hasil tambahannya 2 pesanan. Coba juga tanpa argumen (kedua angka rekap nol) dan `1 5 6` (6 unit, 1 ditolak).

[Roadmap](../README.md) · [Bab sebelumnya](../bab-08/README.md) · [Bab berikutnya](../bab-10/README.md)
