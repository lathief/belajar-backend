# Bab 01 — Program pertama

Cargo membaca `Cargo.toml` untuk menemukan program yang akan dibangun. Eksekusi masuk melalui `fn main()`. Tanda `!` pada `println!` menandai makro. Akhiri perintah cetak dengan `;`.

## Baca dan jalankan

Buka [main.rs](main.rs). Dari direktori `learn-rust/bab-01`, jalankan perintah berikut.

```sh
cargo run --quiet --bin bab-01
```

Keluaran program yang diharapkan

```text
layanan pesanan siap
```

## Telusuri program

Baca `fn main()` sebagai fungsi tanpa parameter. Kurung kurawal menandai tubuhnya. Baris yang diawali `//` dilewati compiler. Teks di antara tanda petik adalah string literal.

## Latihan mandiri

Tambahkan baris kedua yang mencetak `mode lokal`. Jalankan ulang dan periksa urutan kedua baris.

[Roadmap](../README.md) · [Bab berikutnya](../bab-02/README.md)
