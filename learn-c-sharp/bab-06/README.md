# Bab 06 — Record, enum, dan switch

`record` membentuk tipe data dengan property dari parameter deklarasinya. `with` membuat salinan dengan perubahan tertentu. `enum` memberi nama pada status. Switch expression menerjemahkan nilai status ke label tampilan.

## Baca dan jalankan

Buka [Program.cs](Program.cs). Dari direktori `learn-c-sharp/bab-06`, jalankan perintah berikut.

```sh
dotnet run --file Program.cs
```

Keluaran program yang diharapkan

```text
PS-1 lunas
```

## Telusuri program

`awal` mempunyai status `Baru`. Ekspresi `with` menyalin kode pesanan ke `akhir` dan mengganti status salinan dengan `Bayar`, yang kemudian dipetakan ke teks `lunas`.

## Latihan mandiri

Tambahkan status `Batal` dan label `dibatalkan`. Cetak juga status pesanan awal untuk membuktikan `with` menyimpan nilai awal.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-05/README.md) · [Bab berikutnya](../bab-07/README.md)
