# Bab 08 — Parsing, error, dan validasi

Fungsi Go dapat mengembalikan lebih dari satu nilai. Pola `(int, error)` membawa hasil beserta status kesalahan. `strconv.Atoi` mencoba membaca teks sebagai bilangan bulat desimal. Periksa `err != nil` sebelum memakai angkanya.

## Baca dan jalankan

Buka [main.go](main.go). Dari direktori `learn-go/bab-08`, jalankan perintah berikut.

```sh
go run main.go
```

Keluaran program yang diharapkan

```text
jumlah=2
di luar batas
bukan bilangan bulat
```

## Telusuri program

Input `0` berhasil dikonversi tetapi berada di luar rentang. Input `x` sudah gagal saat konversi. Pada kedua kasus, `continue` membawa loop ke input berikutnya. Nilai `nil` dikembalikan saat konversi dan validasi berhasil.

## Latihan mandiri

Tambahkan `-1`, `5`, `6`, dan `2.5`. Prediksi pesan masing-masing. Hanya `5` yang menghasilkan jumlah valid, sedangkan `2.5` harus gagal saat konversi.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-07/README.md) · [Bab berikutnya](../bab-09/README.md)
