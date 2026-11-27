# Bab 16 — Retry dengan backoff dan deadline

Retry dapat membantu saat operasi gagal sementara, tetapi harus berhenti saat deadline tercapai. Contoh memakai jam simulasi agar hasilnya deterministik dan tidak perlu menunggu.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
percobaan=1 gagal
percobaan=2 gagal
deadline_tercapai
```

## Telusuri program

Jeda backoff bertambah dari 2 menjadi 4 satuan waktu. Percobaan berikutnya tidak dimulai karena jeda akan melewati deadline.

## Latihan mandiri

Naikkan deadline menjadi 10 dan lihat berapa banyak percobaan yang dapat dimulai.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-15/README.md) · [Bab berikutnya](../bab-17/README.md)
