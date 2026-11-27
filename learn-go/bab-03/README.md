# Bab 03 — Kondisi dan perulangan

`if` menerima kondisi boolean. Operator `&&` mensyaratkan dua kondisi benar. `for ... range` mengunjungi isi slice secara berurutan. Dalam `for _, n := range daftar`, `_` mengabaikan indeks dan `n` memegang nilai elemen.

## Baca dan jalankan

Buka [main.go](main.go). Dari direktori `learn-go/bab-03`, jalankan perintah berikut.

```sh
go run main.go
```

Keluaran program yang diharapkan

```text
valid=2
```

## Telusuri program

Nilai 0 gagal pada batas bawah dan 8 gagal pada batas atas. Hanya 1 dan 4 yang menjalankan `valid++`. Blok `if` dan `for` dibatasi kurung kurawal.

## Latihan mandiri

Tambahkan 5 dan 6 sehingga hasil menjadi 3. Kemudian tulis ulang loop memakai `for i := 0; i < len(daftar); i++` dan baca `daftar[i]`. Pastikan hasil tetap sama.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-02/README.md) · [Bab berikutnya](../bab-04/README.md)
