# Belajar Docker untuk backend developer

Dalam sembilan bab, kamu akan mengubah katalog buku dari proses sekali jalan menjadi layanan HTTP yang bisa diperbarui dan dipulihkan. Setiap bab berisi praktik, hasil yang bisa diamati, latihan mandiri, dan perintah untuk membereskan resource latihan.

## Persiapan

Gunakan PowerShell dan Docker Desktop dengan Linux containers. Siapkan koneksi internet untuk mengambil image Alpine, BusyBox, dan Go pada pemakaian pertama. Compiler Go dijalankan di container.

```powershell
docker version
docker compose version
```

Pastikan `docker version` menampilkan bagian `Client` dan `Server`. Jika bagian server gagal diakses, jalankan Docker Desktop dan tunggu sampai engine siap. Mulai dari folder `Projects` dan masuk ke bab pertama.

```powershell
cd learn-docker/bab-01
```

Jalankan perintah setiap bab dari folder bab tersebut. Untuk berpindah dari Bab 01 ke Bab 02, gunakan `cd ../bab-02`. Bab 09 memakai kode sumber Bab 08 lewat path relatif, jadi pertahankan keduanya dalam folder `learn-docker`.

## Beginner

| Bab | Materi | Hasil praktik |
| --- | --- | --- |
| 01 | [Image dan container](bab-01/README.md) | Menjalankan proses satu kali dan membaca status akhirnya |
| 02 | [Port dan siklus hidup server](bab-02/README.md) | Mengakses HTTP dari host, menghentikan, dan menyalakan server kembali |
| 03 | [Membangun image dengan Dockerfile](bab-03/README.md) | Membangun image yang memuat halaman katalog |
| 04 | [Mengedit halaman dengan bind mount](bab-04/README.md) | Mengubah stok di editor dan melihat hasilnya lewat HTTP |
| 05 | [Menyimpan data dalam named volume](bab-05/README.md) | Membaca data setelah container penulis dihapus |

## Intermediate

| Bab | Materi | Hasil praktik |
| --- | --- | --- |
| 06 | [Jaringan antar-container](bab-06/README.md) | Menghubungi server melalui nama container pada jaringan bridge |
| 07 | [Menjalankan dua layanan dengan Compose](bab-07/README.md) | Mendefinisikan server dan klien dalam satu berkas |
| 08 | [Multi-stage build untuk layanan Go](bab-08/README.md) | Memisahkan compiler Go dari image yang menjalankan aplikasi |
| 09 | [Proyek akhir — katalog yang siap dibaca](bab-09/README.md) | Menguji kesiapan, persistensi, dan pemulihan layanan |
