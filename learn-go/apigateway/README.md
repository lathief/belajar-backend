# API gateway

Gateway menerima `GET /api/products?id=P01`, menghapus awalan `/api`, lalu meneruskan permintaan beserta query ke backend lokal. Ketiga versi memakai `httputil.ReverseProxy` dan mengembalikan status serta isi respons backend.

Dari folder `learn-go/apigateway`, gunakan Go 1.25 atau lebih baru. Jalankan backend di terminal pertama.

```powershell
go run ./backend
```

Di terminal kedua, dari folder yang sama, pilih satu gateway.

| Server | Perintah | Port |
| --- | --- | --- |
| Fiber | `go run ./fiber` | 8081 |
| Gin | `go run ./gin` | 8082 |
| net/http | `go run ./plain` | 8083 |

Di terminal ketiga, coba versi Fiber. Ganti port untuk Gin atau net/http.

```powershell
curl.exe -i "http://127.0.0.1:8081/api/products?id=P01"
curl.exe -i "http://127.0.0.1:8081/api/products?id=P99"
```

Produk `P01` menghasilkan `200 OK` dengan nama `Buku Go`. Produk `P99` menghasilkan `404 Not Found` dari backend. Hentikan backend di port 9000 dan ulangi permintaan untuk melihat `502 Bad Gateway`.
