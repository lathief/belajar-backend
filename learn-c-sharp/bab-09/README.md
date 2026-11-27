# Bab 09 — Proyek mini rekapan pesanan

Program top-level menerima argumen melalui `args`. Setelah `TryParse` sukses, pola `is >= 1 and <= 5` memeriksa kedua batas. Pisahkan hitungan unit dari jumlah pesanan yang ditolak agar data buruk tetap terlihat pada laporan.

## Baca dan jalankan

Buka [Program.cs](Program.cs). Dari direktori `learn-c-sharp/bab-09`, jalankan perintah berikut.

```sh
dotnet run --file Program.cs -- 2 0 x 3
```

Keluaran program yang diharapkan

```text
unit=5 ditolak=2
```

## Telusuri program

Untuk argumen `x`, evaluasi berhenti setelah `TryParse` menghasilkan `false`. Untuk `0`, pola rentang yang gagal. Kedua jalur masuk ke `else` yang sama.

## Latihan mandiri

Tambahkan hitungan pesanan valid. Untuk `2 0 x 3` hasil tambahannya 2 pesanan. Tanpa argumen kedua angka nol. Untuk `1 5 6` harapkan 6 unit dan 1 ditolak.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-08/README.md) · [Bab berikutnya](../bab-10/README.md)
