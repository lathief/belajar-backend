# Bab 07 — Menjalankan dua layanan dengan Compose

[compose.yaml](compose.yaml) mendeskripsikan server `katalog` dan klien sekali jalan `pembaca`. Compose menyiapkan jaringan untuk proyek. Nama layanan `katalog` dapat dipakai klien sebagai nama host.

## Jalankan katalog dan baca isinya

```powershell
docker compose config
docker compose up -d katalog
curl.exe -s http://127.0.0.1:8077/buku.txt
docker compose run --rm pembaca
docker compose ps
```

Curl dari komputer memakai port 8077. Klien di jaringan Compose memakai `katalog:8080`. Keduanya membaca isi [buku.txt](katalog/buku.txt) dan mencetak `Buku Go tersedia.`. Setelah klien selesai, `compose ps` menampilkan server yang masih berjalan.

## Latihan

Ubah port host dari `8077` menjadi `8177` di `compose.yaml`. Terapkan perubahan dengan `docker compose up -d katalog`, lalu coba curl ke port baru dan jalankan pembaca lagi. Mengapa URL di konfigurasi pembaca tetap berfungsi?

## Bereskan

```powershell
docker compose down
```

Kembalikan port ke 8077 jika mengerjakan latihan.

[Bab sebelumnya](../bab-06/README.md) · [Roadmap](../README.md) · [Bab berikutnya](../bab-08/README.md)
