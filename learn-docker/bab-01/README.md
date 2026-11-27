# Bab 01 — Image dan container

Image berisi program beserta berkas yang dibutuhkan untuk menjalankannya. Saat membuat container dari image Alpine, kita bisa memilih perintah Linux yang menjadi proses utamanya.

## Jalankan proses pertama

```powershell
docker pull alpine:3.22
docker image ls alpine:3.22
docker run --name ld01-katalog alpine:3.22 echo Katalog siap
docker ps -a --filter name=ld01-katalog
docker logs ld01-katalog
```

Baris `docker run` mencetak `Katalog siap`. Setelah `echo` selesai, status container menjadi `Exited (0)`. Angka nol berarti proses berakhir dengan sukses. Teksnya masih bisa dibaca lewat `docker logs`.

## Latihan

Jalankan `docker run --rm alpine:3.22 cat /etc/alpine-release`. Amati versi Alpine yang tercetak. Opsi `--rm` menghapus container latihan kedua begitu perintah selesai.

## Bereskan

```powershell
docker rm ld01-katalog
```

Image Alpine tetap tersimpan untuk bab berikutnya.

[Roadmap](../README.md) · [Bab berikutnya](../bab-02/README.md)
