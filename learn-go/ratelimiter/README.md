# Rate limiter

Endpoint `GET /ping` membalas `pong`. Setiap proses server punya satu kuota yang dipakai bersama oleh semua klien. Permintaan pertama membuka jendela 10 detik. Tiga permintaan pertama dalam jendela itu lolos. Permintaan berikutnya mendapat `429 Too Many Requests` dan header `Retry-After` berisi waktu tunggu dalam detik.

Dari folder `learn-go/ratelimiter`, gunakan Go 1.25 atau lebih baru dan pilih satu server.

| Server | Perintah | Port |
| --- | --- | --- |
| Fiber | `go run ./fiber` | 8081 |
| Gin | `go run ./gin` | 8082 |
| net/http | `go run ./plain` | 8083 |

Jalankan perintah ini di terminal PowerShell lain. Ganti port sesuai server yang dipilih.

```powershell
1..4 | ForEach-Object { curl.exe -i http://127.0.0.1:8081/ping }
Start-Sleep -Seconds 11
curl.exe -i http://127.0.0.1:8081/ping
```

Dengan kuota awal penuh, urutan statusnya `200, 200, 200, 429`, lalu `200` setelah menunggu. Fiber memakai middleware `limiter`. Gin dan net/http memakai penghitung di `internal/limit` dengan mutex agar permintaan bersamaan tetap mengikuti kuota.
