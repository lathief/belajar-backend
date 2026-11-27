# Bab 06 — Jaringan antar-container

Sebuah layanan bisa diakses oleh container lain lewat nama pada jaringan bridge buatan sendiri. Di sini server katalog hanya perlu menerima koneksi dari container pembaca.

## Hubungkan server dan pembaca

```powershell
docker network create ld06-jaringan
docker run -d --name ld06-katalog --network ld06-jaringan busybox:1.37 sh -c 'mkdir -p /srv && echo Buku-Go > /srv/buku.txt && exec httpd -f -p 8080 -h /srv'
docker run --rm --network ld06-jaringan alpine:3.22 wget -qO- http://ld06-katalog:8080/buku.txt
```

Pembaca mencetak `Buku-Go`. Docker menerjemahkan nama `ld06-katalog` ke alamat container pada jaringan itu. Koneksi memakai port internal 8080 sehingga server bisa dihubungi tanpa memublikasikan port ke komputer.

## Latihan

Periksa `docker network inspect ld06-jaringan` dan temukan nama server pada bagian `Containers`. Ulangi perintah pembaca dengan URL `http://127.0.0.1:8080/buku.txt`. Permintaan ini gagal karena alamat loopback tersebut mengarah ke container pembaca sendiri.

## Bereskan

```powershell
docker stop ld06-katalog
docker rm ld06-katalog
docker network rm ld06-jaringan
```

[Bab sebelumnya](../bab-05/README.md) · [Roadmap](../README.md) · [Bab berikutnya](../bab-07/README.md)
