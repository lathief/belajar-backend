# Basic auth

Endpoint `GET /profil` membaca kredensial dari header HTTP `Authorization`. Akun latihan lokal menggunakan nama pengguna `demo` dan kata sandi `belajar`.

Dari folder `learn-go/basicauth`, gunakan Go 1.25 atau lebih baru dan pilih satu server.

| Server | Perintah | Port |
| --- | --- | --- |
| Fiber | `go run ./fiber` | 8081 |
| Gin | `go run ./gin` | 8082 |
| net/http | `go run ./plain` | 8083 |

Coba tiga permintaan dari terminal PowerShell lain. Ganti port sesuai server yang dipilih.

```powershell
curl.exe -i http://127.0.0.1:8081/profil
curl.exe -i -u demo:salah http://127.0.0.1:8081/profil
curl.exe -i -u demo:belajar http://127.0.0.1:8081/profil
```

Dua permintaan pertama menghasilkan `401 Unauthorized` dengan header `WWW-Authenticate`. Permintaan terakhir menghasilkan `200 OK` dengan JSON `{"user":"demo"}`. Fiber dan Gin memakai middleware basic auth bawaan masing-masing. Versi net/http membaca kredensial dengan `r.BasicAuth()`.
