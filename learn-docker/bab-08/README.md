# Bab 08 — Multi-stage build untuk layanan Go

Pada [Dockerfile](Dockerfile), tahap `build` mengompilasi [server Go](main.go). Tahap terakhir menyalin binary dengan `COPY --from=build` dan menyertakan [data katalog](data/buku.txt). Compiler cukup tersedia pada tahap build. Server di image akhir berjalan dengan UID 10001. [Dokumentasi multi-stage build](https://docs.docker.com/build/building/multi-stage/)

## Bangun kedua tahap

```powershell
docker build --target build -t ld08-katalog:builder .
docker build -t ld08-katalog:runtime .
docker image ls ld08-katalog
docker run -d --name ld08-api -p 127.0.0.1:8078:8080 ld08-katalog:runtime
curl.exe -s http://127.0.0.1:8078/buku
```

Endpoint `/buku` membalas `Buku Go tersedia.`. Bandingkan kolom ukuran kedua tag. Tag `builder` mencakup toolchain Go, sedangkan tag `runtime` berisi Alpine, binary, dan berkas data.

## Latihan

```powershell
docker run --rm ld08-katalog:builder go version
docker exec ld08-api id -u
docker logs ld08-api
```

Temukan versi compiler, UID `10001`, dan pesan startup dari server. Ubah isi `data/buku.txt`, lalu prediksi tahap mana yang harus dikerjakan ulang saat build berikutnya. Periksa prediksi melalui `docker build --progress=plain -t ld08-katalog:runtime .`.

## Bereskan

```powershell
docker stop ld08-api
docker rm ld08-api
docker image rm ld08-katalog:builder ld08-katalog:runtime
```

Jika kamu mengubah isi data, kembalikan ke `Buku Go tersedia.` agar proyek akhir dimulai dengan contoh yang sama.

[Bab sebelumnya](../bab-07/README.md) · [Roadmap](../README.md) · [Bab berikutnya](../bab-09/README.md)
