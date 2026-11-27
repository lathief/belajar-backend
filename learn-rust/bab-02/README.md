# Bab 02 — Nilai, tipe, dan operator

`let` mengikat nama ke nilai. Tambahkan `mut` bila nilainya akan diperbarui. `u32` menyimpan bilangan bulat tanpa tanda, cocok untuk jumlah kecil pada contoh ini. Operator `*` menghitung subtotal sebelum nilainya disisipkan ke teks.

## Baca dan jalankan

Buka [main.rs](main.rs). Dari direktori `learn-rust/bab-02`, jalankan perintah berikut.

```sh
cargo run --quiet --bin bab-02
```

Keluaran program yang diharapkan

```text
subtotal=15000
```

## Telusuri program

Setelah `jumlah += 1`, nilainya 3. Saat Rust mengalikan `harga` dengan `jumlah`, tipe `jumlah` dapat disimpulkan dari `harga`. Menghapus `mut` menyebabkan assignment kedua ditolak.

## Latihan mandiri

Ubah jumlah awal menjadi 3. Dengan tambahan satu barang, subtotal harus menjadi 20000.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-01/README.md) · [Bab berikutnya](../bab-03/README.md)
