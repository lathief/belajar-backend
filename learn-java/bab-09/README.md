# Bab 09 — Proyek mini rekapan pesanan

`args` berisi argumen sesudah nama program. Method `jumlahValid` mengubah satu teks menjadi jumlah 1–5 atau nilai penanda 0. Karena pesanan sah selalu positif, pemanggil dapat menggunakan penanda ini saat menyusun rekap.

## Baca dan jalankan

Buka [Main.java](Main.java). Dari direktori `learn-java/bab-09`, jalankan perintah berikut.

```sh
java Main.java 2 0 x 3
```

Keluaran program yang diharapkan

```text
unit=5 ditolak=2
```

## Telusuri program

`jumlahValid` dipanggil sekali per argumen. Hasil 2 dan 3 menambah unit. Kedua nilai penanda 0 hanya menambah hitungan entri ditolak.

## Latihan mandiri

Tambahkan hitungan pesanan valid. Untuk `2 0 x 3` harus ada 2 pesanan valid. Tanpa argumen, kedua angka rekap nol. Untuk `1 5 6` harapkan 6 unit dan 1 ditolak.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-08/README.md) · [Bab berikutnya](../bab-10/README.md)
