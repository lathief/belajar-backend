# Bab 01 — Program pertama

`package main` menandai paket yang dapat dibangun menjadi program. Ketika program berjalan, Go memanggil `func main()`. Deklarasi `import` menyediakan paket `fmt` untuk menulis keluaran ke terminal.

## Baca dan jalankan

Buka [main.go](main.go). Dari direktori `learn-go/bab-01`, jalankan perintah berikut.

```sh
go run main.go
```

Keluaran program yang diharapkan

```text
layanan pesanan siap
```

## Telusuri program

`fmt.Println` menerima string di dalam kurung dan menambahkan baris baru. Teks setelah `//` adalah komentar untuk pembaca kode.

## Latihan mandiri

Cetak `mode lokal` di baris kedua. Balik urutan dua pemanggilan `Println` dan prediksi output sebelum menjalankannya.

[Roadmap](../README.md) · [Bab berikutnya](../bab-02/README.md)
