# Bab 08 — Parsing, exception, dan validasi

Teks dari pengguna perlu dikonversi dan divalidasi sebelum menjadi nilai bisnis. `filter_var` mengembalikan `false` untuk format yang tidak valid; exception digunakan untuk melaporkan input yang tidak memenuhi aturan.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
2
input_tidak_valid
jumlah_di_luar_batas
```

## Telusuri program

`parseJumlah` membedakan teks yang bukan integer dari integer di luar rentang 1–5. `try/catch` menangani kedua kegagalan secara eksplisit.

## Latihan mandiri

Uji `"5"`, `"0"`, dan `"dua"`. Catat mana yang valid, melanggar batas, atau gagal parsing.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-07/README.md) · [Bab berikutnya](../bab-09/README.md)
