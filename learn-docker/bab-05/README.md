# Bab 05 — Menyimpan data dalam named volume

Data yang ditulis ke lapisan container ikut hilang saat container dihapus. Untuk menyimpan catatan stok lintas container, pasang volume yang dikelola Docker ke direktori `/data`.

## Tulis lalu baca dari container lain

```powershell
docker volume create ld05-stok
docker run --rm --mount type=volume,source=ld05-stok,target=/data alpine:3.22 sh -c 'echo stok=3 > /data/stok.txt'
docker run --rm --mount type=volume,source=ld05-stok,target=/data,readonly alpine:3.22 cat /data/stok.txt
```

Container penulis sudah dihapus oleh `--rm` saat container pembaca mulai. Keluaran `stok=3` menunjukkan bahwa keduanya mengakses volume yang sama.

## Latihan

Ulangi perintah penulis dengan `stok=7`, lalu jalankan pembaca lagi. Gunakan `docker volume inspect ld05-stok` untuk melihat nama dan driver volume.

## Bereskan

Perintah ini menghapus volume latihan beserta isinya setelah kamu selesai membaca hasilnya.

```powershell
docker volume rm ld05-stok
```

[Bab sebelumnya](../bab-04/README.md) · [Roadmap](../README.md) · [Bab berikutnya](../bab-06/README.md)
