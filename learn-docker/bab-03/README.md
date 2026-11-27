# Bab 03 — Membangun image dengan Dockerfile

Di bab ini halaman katalog menjadi bagian dari image. Buka [Dockerfile](Dockerfile) dan [index.html](index.html). `FROM` memilih image dasar, `WORKDIR` menetapkan direktori kerja, `COPY` memasukkan halaman, dan `CMD` menentukan perintah saat container dijalankan.

## Bangun dan jalankan

```powershell
docker build -t ld03-katalog:v1 .
docker run -d --name ld03-web -p 127.0.0.1:8073:8080 ld03-katalog:v1
curl.exe -s http://127.0.0.1:8073/
```

Titik pada akhir perintah build memilih folder bab ini sebagai *build context*, yaitu sumber berkas yang dapat disalin oleh Docker. Respons memuat `Buku Go tersedia.`. `EXPOSE 8080` mendokumentasikan port aplikasi. Opsi `-p` memublikasikannya ke komputer.

## Latihan

Ubah nama buku di `index.html`, lalu panggil URL yang sama. Container pertama masih menyajikan salinan lama. Bangun tag kedua dan bandingkan hasilnya.

```powershell
docker build -t ld03-katalog:v2 .
docker run --rm ld03-katalog:v2 cat /srv/index.html
```

## Bereskan

```powershell
docker stop ld03-web
docker rm ld03-web
docker image rm ld03-katalog:v1
```

Jika mengerjakan latihan tag kedua, hapus juga dengan `docker image rm ld03-katalog:v2`.

[Bab sebelumnya](../bab-02/README.md) · [Roadmap](../README.md) · [Bab berikutnya](../bab-04/README.md)
