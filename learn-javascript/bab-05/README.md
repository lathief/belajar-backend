# Bab 05 — Objek, destructuring, dan penyalinan

Objek menyimpan pasangan nama dan nilai. Destructuring mengambil properti berdasarkan nama. Spread `...` menyalin properti lapisan pertama. Properti setelah spread menimpa nilai yang telah disalin. Ini berguna untuk mengubah status sambil menyimpan keadaan awal.

## Baca dan jalankan

Buka [main.mjs](main.mjs). Dari direktori `learn-javascript/bab-05`, jalankan perintah berikut.

```sh
node main.mjs
```

Keluaran program yang diharapkan

```text
PS-1 baru -> lunas
```

## Telusuri program

Setelah spread, `akhir` memiliki properti baru hasil penyalinan. Pada contoh ini semua nilai properti bersifat primitif. Jika properti berisi objek lain, salinan dangkal masih mengacu pada objek bersarang yang sama.

## Latihan mandiri

Tambahkan properti `jumlah` bernilai 2. Pada objek baru, ubah menjadi 3. Cetak kedua jumlah dan pastikan objek awal masih memuat 2.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-04/README.md) · [Bab berikutnya](../bab-06/README.md)
