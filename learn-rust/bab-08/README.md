# Bab 08 — Result, parsing, dan operator ?

`parse` dapat gagal ketika teks bukan angka. `Result<T, E>` membawa nilai sukses atau alasan gagal. `map_err` menyamakan jenis error, sedangkan `?` segera mengembalikan error itu ke pemanggil. Validasi domain masih diperlukan sesudah parsing.

## Baca dan jalankan

Buka [main.rs](main.rs). Dari direktori `learn-rust/bab-08`, jalankan perintah berikut.

```sh
cargo run --quiet --bin bab-08
```

Keluaran program yang diharapkan

```text
jumlah=2
harus positif
bukan angka
```

## Telusuri program

Untuk `x`, `parse` menghasilkan `Err`. Operator `?` menghentikan fungsi sebelum pemeriksaan rentang. Untuk `0`, parsing berhasil dan pemeriksaan bisnis yang menghasilkan `Err`.

## Latihan mandiri

Tambahkan batas maksimum 5. Uji `5`, `6`, dan `-1`. Hanya input pertama yang boleh berhasil.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-07/README.md) · [Bab berikutnya](../bab-09/README.md)
