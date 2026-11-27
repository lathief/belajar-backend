# Bab 11 — Transisi status pesanan

## Studi kasus

Petugas gudang mencoba mengirim pesanan sebelum pembayaran tercatat. Anda akan memisahkan aksi yang diminta petugas dari status yang boleh disimpan.

**Target belajar** — Menolak aksi yang tidak cocok dengan status saat ini.

## Aturan kasus

| Status awal | Aksi | Tujuan |
| --- | --- | --- |
| `baru` | `bayar` | `lunas` |
| `baru` | `batal` | `batal` |
| `lunas` | `kirim` | `dikirim` |

Pasangan lain ditolak dan mempertahankan status awal. `dikirim` dan `batal` adalah status terminal.

## Baca dan jalankan

Buka [main.go](main.go). Dari direktori `learn-go/bab-11`, jalankan contoh berikut.

```sh
go run main.go
```

Keluaran yang diharapkan

```text
baru kirim -> baru ok=false
baru bayar -> lunas ok=true
lunas batal -> lunas ok=false
lunas kirim -> dikirim ok=true
dikirim bayar -> dikirim ok=false
```

## Telusuri program

Tipe bernama `status` mewakili keadaan pesanan. Dua nilai balik memisahkan kandidat status dari keputusan menerima aksi.

Aksi `kirim` pertama ditolak sehingga `bayar` masih dimulai dari `baru`. Sesudah pembayaran, pembatalan ditolak. Pengiriman berikutnya berhasil karena pesanan sudah `lunas`. Berbeda dengan kunci permintaan pada Bab 10, hasil di sini bergantung pada status ketika aksi diproses.

## Latihan mandiri

1. Mulai dengan pesanan baru dan jalankan `batal`, `bayar`, `kirim`. Hanya aksi pertama boleh berhasil dan status akhir harus `batal`.
2. Tambahkan jalur dari `dikirim` ke `selesai` lewat aksi `terima`. Pastikan `terima` ditolak untuk pesanan baru dan ditolak jika diulangi pada pesanan selesai.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-10/README.md) · [Bab berikutnya](../bab-12/README.md)
