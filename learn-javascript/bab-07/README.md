# Bab 07 — Modul dan validasi input

`export` memberi nama yang dapat diimpor modul lain. Impor lokal menyertakan jalur relatif dan ekstensi `.mjs`. Validator mengembalikan angka valid atau `null`, sehingga pemanggil bisa memilih respons untuk input buruk.

## Baca dan jalankan

Buka [main.mjs](main.mjs) dan [modul validasi](validasi.mjs). Dari direktori `learn-javascript/bab-07`, jalankan perintah berikut.

```sh
node main.mjs
```

Keluaran program yang diharapkan

```text
2
ditolak
```

## Telusuri program

Di modul validator, pola `^` sampai `$` memeriksa seluruh input dan `\d+` meminta satu digit atau lebih. `2` lolos pemeriksaan rentang. `x` menghasilkan `null` sebelum konversi.

## Latihan mandiri

Uji `0`, `2.5`, dan string kosong. Ketiganya harus ditolak. Ubah batas maksimum menjadi 10 dan buktikan 10 diterima sedangkan 11 ditolak.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-06/README.md) · [Bab berikutnya](../bab-08/README.md)
