# Bab 03 — Kondisi dan perulangan

`for` mengambil satu nilai per iterasi. `if` menerima ekspresi boolean. Operator `&&` meminta kedua syarat benar. Hitung hanya jumlah pesanan yang berada di antara batas bawah dan batas atas.

## Baca dan jalankan

Buka [main.rs](main.rs). Dari direktori `learn-rust/bab-03`, jalankan perintah berikut.

```sh
cargo run --quiet --bin bab-03
```

Keluaran program yang diharapkan

```text
valid=2
```

## Telusuri program

Counter berawal dari 0. Input 0 gagal pada batas bawah, 1 dan 4 masing-masing menambah counter, lalu 8 gagal pada batas atas.

## Latihan mandiri

Tambahkan 5 dan 6 ke input. Hasil harus menjadi 3. Jelaskan mengapa hanya satu nilai tambahan yang lolos.

Coba tulis kembali loop dengan `let mut i = 0;` dan `while i < daftar.len()`. Simpan input lebih dulu sebagai `let daftar = [0, 1, 4, 8];`, ambil nilai lewat `daftar[i]`, lalu naikkan `i` di setiap iterasi. Periksa bahwa hasil dengan input awal tetap 2. Apa yang terjadi jika indeks berhenti bertambah?

[Roadmap](../README.md) · [Bab sebelumnya](../bab-02/README.md) · [Bab berikutnya](../bab-04/README.md)
