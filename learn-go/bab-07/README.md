# Bab 07 — Interface dan aturan ongkir

Interface menyebutkan method yang diperlukan pemanggil. Tipe yang memiliki method dengan nama, parameter, dan hasil yang cocok otomatis memenuhi interface tersebut. Di sini fungsi `tagihan` bisa memakai dua aturan ongkir.

## Baca dan jalankan

Buka [main.go](main.go). Dari direktori `learn-go/bab-07`, jalankan perintah berikut.

```sh
go run main.go
```

Keluaran program yang diharapkan

```text
kirim=17000
ambil=15000
```

## Telusuri program

Dengan subtotal 15000, `TarifTetap` menambah 2000 sedangkan `AmbilSendiri` menghasilkan biaya 0. Keduanya diberikan ke fungsi yang sama. Struct kosong `AmbilSendiri{}` cukup karena aturan itu tidak perlu menyimpan data.

## Latihan mandiri

Tambahkan tipe `GratisMulai` dengan field `Minimum` dan `Biaya`. Untuk minimum 20000 dan biaya 2000, subtotal 15000 harus ditagih 17000 dan subtotal 20000 tetap 20000. Pertahankan fungsi `tagihan` yang ada.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-06/README.md) · [Bab berikutnya](../bab-08/README.md)
