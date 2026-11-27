# Bab 03 — Kondisi dan iterasi

`int[]` membatasi isi array ke bilangan bulat. `foreach` memberi akses ke satu elemen pada setiap iterasi. `&&` meminta kedua syarat benar. Gunakan blok `if` untuk mengubah hitungan hanya ketika jumlah berada di dalam batas.

## Baca dan jalankan

Buka [Program.cs](Program.cs). Dari direktori `learn-c-sharp/bab-03`, jalankan perintah berikut.

```sh
dotnet run --file Program.cs
```

Keluaran program yang diharapkan

```text
valid=2
```

## Telusuri program

Variabel `n` berisi nilai saat ini. Nilai 0 melewati blok penambahan, 1 dan 4 masing-masing memasukinya, lalu 8 melewatinya lagi.

## Latihan mandiri

Tambahkan 5 dan 6. Hasil harus menjadi 3. Tambahkan cabang `else` untuk mencetak nilai yang ditolak.

Coba `int i = 0;` dan `while (i < daftar.Length)`. Baca nilai lewat `daftar[i]` dan naikkan `i` di setiap iterasi. Periksa bahwa hasil dengan input awal tetap 2. Apa yang terjadi jika indeks berhenti bertambah?

[Roadmap](../README.md) · [Bab sebelumnya](../bab-02/README.md) · [Bab berikutnya](../bab-04/README.md)
