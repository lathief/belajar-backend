# Bab 06 — Struct, enum, dan match

`struct` mengelompokkan field bernama. `enum` membatasi status yang mungkin. `match` mewajibkan setiap varian tercakup sehingga penambahan status baru terlihat saat kompilasi.

## Baca dan jalankan

Buka [main.rs](main.rs). Dari direktori `learn-rust/bab-06`, jalankan perintah berikut.

```sh
cargo run --quiet --bin bab-06
```

Keluaran program yang diharapkan

```text
PS-1 menunggu
PS-2 lunas
```

## Telusuri program

`Pesanan` menyatukan kode dengan status. Pada iterasi kedua, pola `Status::Bayar` cocok dan seluruh ekspresi `match` menghasilkan teks `lunas`.

## Latihan mandiri

Tambahkan varian `Batal` dan satu pesanan untuk varian itu. Pola baru harus mencetak `dibatalkan`.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-05/README.md) · [Bab berikutnya](../bab-07/README.md)
