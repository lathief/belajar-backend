# Bab 09 — Proyek akhir — katalog yang siap dibaca

Jalankan layanan Go, tunggu sampai katalog bisa dibaca, lalu buktikan bahwa perubahan stok bertahan saat container diganti. [compose.yaml](compose.yaml) membangun [aplikasi Bab 08](../bab-08/Dockerfile) langsung dari sumbernya. Volume `stok` dipasang di `/srv/data`. Saat pertama dipasang dalam keadaan kosong, Docker menyalin data awal dari image ke volume. [Dokumentasi volume](https://docs.docker.com/engine/storage/volumes/)

## Uji kesiapan layanan

```powershell
docker compose up -d --build
docker compose ps -a
docker compose logs pembaca
curl.exe -i http://127.0.0.1:8079/buku
```

Healthcheck meminta `/buku` dari dalam container katalog. Compose menunggu status `healthy` sebelum memulai `pembaca`. Klien membaca satu kali, mencetak isi katalog, lalu berakhir dengan kode 0. Periksa log klien dan respons HTTP `200`. [Urutan startup Compose](https://docs.docker.com/compose/how-tos/startup-order/)

## Ganti container dan periksa data

```powershell
docker compose exec katalog sh -c 'echo stok=7 > /srv/data/buku.txt'
docker compose down
docker compose up -d
docker compose logs pembaca
curl.exe -s http://127.0.0.1:8079/buku
```

Respons yang diharapkan adalah `stok=7`. Perintah `down` tadi menghapus container dan jaringan proyek. Saat dijalankan lagi, server memasang kembali volume yang menyimpan nilai tersebut.

## Latihan diagnosis

Ubah sementara nama berkas data dari dalam container dan amati akibatnya.

```powershell
docker compose exec katalog mv /srv/data/buku.txt /srv/data/buku.bak
curl.exe -i http://127.0.0.1:8079/buku
docker compose logs --tail 5 katalog
docker compose ps
```

HTTP langsung memberi status `503`. Setelah lima pemeriksaan berturut-turut gagal, status kesehatan menjadi `unhealthy`. Jalankan `docker compose ps` lagi setelah sekitar 20 detik. Pulihkan data, tunggu pemeriksaan berikutnya, lalu pastikan server kembali `healthy` dan pembaca mencetak `stok=7`.

```powershell
docker compose exec katalog mv /srv/data/buku.bak /srv/data/buku.txt
$id = docker compose ps -q katalog
for ($i = 0; $i -lt 15; $i++) {
    if ((docker inspect --format '{{.State.Health.Status}}' $id) -eq 'healthy') { break }
    Start-Sleep -Seconds 1
}
docker compose ps
docker compose run --rm pembaca
```

## Bereskan

Gunakan `docker compose down` jika ingin melanjutkan dengan stok terakhir. Jika latihan selesai dan data stok boleh dihapus, jalankan perintah berikut dari folder bab ini.

```powershell
docker compose down --volumes
```

[Bab sebelumnya](../bab-08/README.md) · [Roadmap](../README.md)
