# Bab 15 — Token bucket untuk permintaan ekspor

**Tingkat intermediate** · Membedakan kapasitas lonjakan dari laju pengisian.

## Studi kasus

Pada Bab 14, petugas bisa membaca ekspor per halaman. Sekarang satu klien mengirim tiga permintaan ekspor pada detik yang sama. Gudang mampu melayani dua sekaligus dan menerima satu tambahan setiap detik. Anda akan menentukan permintaan yang boleh dimulai.

## Aturan kasus

Satu klien memakai satu ember berkapasitas dua token, awalnya penuh. Satu permintaan memerlukan satu token. Sebelum memutuskan, tambahkan satu token per detik berlalu dan batasi jumlahnya pada kapasitas. Jam simulasi berjalan maju dalam detik bulat. `tunda` menyerahkan keputusan mencoba kembali kepada pemanggil.

## Baca dan jalankan

Buka [main.rs](main.rs). Dari direktori `learn-rust/bab-15`, jalankan perintah berikut.

```sh
cargo run --quiet --bin bab-15
```

Keluaran yang diharapkan

```text
t=0 lolos token=1
t=0 lolos token=0
t=0 tunda token=0
t=1 lolos token=0
t=6 lolos token=1
t=6 lolos token=0
t=6 tunda token=0
```

## Telusuri program

`&mut self` membatasi mutasi ember selama satu pemanggilan. `saturating_add` dan `min` menjaga penambahan serta kapasitas token.

Tiga kedatangan pada detik 0 menghabiskan dua token. Mengulang pada detik yang sama tidak mengisi ember. Jeda dari detik 1 ke 6 cukup untuk mengisi penuh, tetapi kapasitas tetap dua. Uji ini memeriksa lonjakan singkat dan jeda panjang sekaligus.

## Latihan mandiri

1. Ubah kapasitas dan jumlah token awal menjadi 3. Di detik 0 ketiga permintaan harus lolos. Hitung kembali sisa token pada detik 1.
2. Tambahkan klien kedua dengan ember tersendiri. Setelah token klien A habis, permintaan pertama klien B di detik yang sama harus tetap lolos. Jelaskan siapa yang menyimpan token jika satu klien dilayani dua instance.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-14/README.md) · [Bab berikutnya](../bab-16/README.md)
