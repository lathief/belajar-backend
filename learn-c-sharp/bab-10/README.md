# Bab 10 — Reservasi stok dan idempotensi

## Studi kasus

Gudang menerima pesanan A untuk tiga unit. Gudang sudah mengurangi stok, tetapi koneksi terputus sebelum klien menerima jawaban. Klien mengirim lagi A dengan jumlah yang sama. Bagaimana gudang bisa memberi jawaban yang sama ketika stok sekarang sudah habis?

Di Bab 09, setiap argumen dihitung sebagai pesanan tersendiri. Di sini, klien memberi setiap permintaan logis sebuah ID dan menggunakan ID itu lagi saat mencoba ulang. Anda akan menyimpan keputusan pertama per ID agar pengiriman berulang hanya berdampak sekali pada stok.

## Aturan reservasi

Contoh memproses satu jenis barang dengan stok awal 5. Semua ID dalam contoh adalah string terisi. Jalankan permintaan berurutan pada satu gudang.

1. Periksa jumlah terlebih dahulu. Jumlah harus berupa bilangan bulat 1–5. Input di luar aturan menghasilkan `jumlah_invalid` dan ID tetap bisa dipakai untuk koreksi.
2. Jika ID sudah tercatat, cocokkan jumlahnya. Jumlah sama mengembalikan keputusan tersimpan. Jumlah valid yang berbeda menghasilkan `konflik` dan bukti awal tetap berlaku.
3. Untuk ID baru, kurangi stok jika mencukupi dan simpan hasil `diterima`. Jika stok kurang, simpan hasil `stok_kurang`. Keputusan bisnis tersebut final untuk ID itu.

Bukti berisi jumlah dan hasil pertama. Kolom `stok` pada keluaran menunjukkan stok gudang **saat baris dicetak**. Pada pengiriman ulang, hasil reservasi bisa sama meskipun angka stok sudah berubah.

## Baca dan jalankan

Buka [Program.cs](Program.cs). Dari direktori `learn-c-sharp/bab-10`, jalankan contoh berikut.

```sh
dotnet run --file Program.cs
```

Keluaran yang diharapkan

```text
A 3 diterima stok=2
B 3 stok_kurang stok=2
C 0 jumlah_invalid stok=2
C 2 diterima stok=0
A 3 diterima stok=0
A 1 konflik stok=0
B 3 stok_kurang stok=0
```

## Telusuri program

`Gudang(int stokAwal)` adalah primary constructor yang memberi nilai awal pada property `Stok`. Setter privat membatasi tempat yang dapat mengubah stok. `TryGetValue` mengembalikan keberadaan kunci dan mengisi `out var awal` dengan record bukti dalam satu pencarian.

Ikuti ID C dari baris ketiga ke baris keempat. Input 0 ditolak sebelum penulisan bukti. Koreksi menjadi 2 masih bisa memakai ID C dan menghabiskan stok. Pada baris kelima, A meminta 3 lagi ketika stok sudah 0. Karena bukti A dibaca sebelum pengecekan stok, program tetap mengembalikan `diterima`. Baris keenam mencoba memakai A untuk jumlah berbeda dan menghasilkan `konflik`.

## Latihan mandiri

1. Tambahkan permintaan A untuk 3 unit sesudah baris terakhir. Harapkan `diterima` dengan stok tetap 0. Kasus ini memeriksa bahwa konflik sebelumnya tidak menimpa bukti A.
2. Tambahkan operasi isi ulang sebanyak 5 setelah batch selesai. Coba lagi B untuk 3 unit, lalu pesan 3 unit memakai ID D. B harus tetap menghasilkan `stok_kurang` dengan stok 5. D harus menghasilkan `diterima` dengan stok 2.
3. Ubah hasil menjadi `enum` dan gunakan switch expression untuk menampilkan label. Pastikan jalur konflik tetap mempertahankan record untuk permintaan pertama.

## Batas studi kasus

Riwayat ada di memori satu proses dan hilang saat proses dimulai ulang. Bayangkan dua instance menerima ID baru yang sama pada saat bersamaan. Keduanya bisa melihat ID itu sebagai ID baru. Untuk layanan dengan database, pengurangan stok dan penyimpanan bukti perlu berada dalam satu transaksi, dengan batas unik pada ID serta pemeriksaan stok yang aman terhadap transaksi bersamaan.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-09/README.md) · [Bab berikutnya](../bab-11/README.md)
