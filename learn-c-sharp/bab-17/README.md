# Bab 17 — Lease worker dan penolakan hasil usang

**Tingkat advanced** · Menguji batas waktu kepemilikan dan token generasi.

## Studi kasus

Pembuatan berkas ekspor dipindahkan ke worker. Worker A melambat dan worker B mengambil alih tugasnya. Hasil A akhirnya datang belakangan. Anda akan menentukan worker mana yang masih berhak menyelesaikan tugas.

## Aturan kasus

Lease berlangsung lima detik dengan rentang aktif `[mulai, kedaluwarsa)`. Klaim sukses menerbitkan token generasi baru. Simpan hanya boleh berhasil pada tugas yang belum selesai, dengan token generasi saat ini dan waktu sebelum kedaluwarsa. Tugas selesai bersifat terminal.

## Baca dan jalankan

Buka [Program.cs](Program.cs). Dari direktori `learn-c-sharp/bab-17`, jalankan perintah berikut.

```sh
dotnet run --file Program.cs
```

Keluaran yang diharapkan

```text
klaim_dini=false
batas=ditolak
lama=ditolak
baru=tersimpan
ulang=ditolak
```

## Telusuri program

`int?` membedakan klaim gagal dari token sukses. Ekspresi `?? throw` pada fixture memastikan klaim yang diharapkan berhasil memang menghasilkan token.

Pada t=4, B belum bisa mengambil alih. Pada t=5, A gagal menyimpan bahkan sebelum B mengklaim. Setelah klaim B, token A juga gagal pada t=6. B menyimpan sekali dan percobaan kedua ditolak. Contoh menjalankan urutan secara serial menggunakan satu jam. Di penyimpanan bersama, terapkan semua predikat finalisasi dan penulisannya secara atomik memakai waktu otoritatif penyimpanan.

## Latihan mandiri

1. Gunakan tugas baru. Simpan dengan token A pada t=4. Hasil harus `tersimpan`, lalu klaim pada t=5 harus gagal karena tugas sudah selesai.
2. Tambahkan perpanjangan lease. Buktikan bahwa token aktif dapat memperpanjang sebelum batas dan token lama gagal memperpanjang setelah pengambilalihan. Pertimbangkan efek samping eksternal jika worker mengirim berkas sebelum finalisasi.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-16/README.md) · [Bab berikutnya](../bab-18/README.md)
