# Contoh worker HTTP di Go

`POST /jobs` memasukkan job ke antrean dan langsung mengembalikan `202 Accepted` dengan respons seperti `{"id":1,"status":"queued"}`. Satu worker di setiap proses mengambil job dan mengirim `GET` ke URL eksternal. Lihat ID job, status HTTP, ukuran respons, atau pesan error di log terminal.

Jalankan dari `test-go/worker` dengan Go 1.25 atau lebih baru.

| Versi | Server | Client keluar | Jalankan | Port |
| --- | --- | --- | --- | --- |
| Fiber | Fiber v2 | Fiber `Agent` | `go run ./fiber` | `8081` |
| Gin | Gin | `http.Client` | `go run ./gin` | `8082` |
| `net/http` | `http.ServeMux` | `http.Client` | `go run ./plain` | `8083` |

Kirim job dari terminal lain.

```powershell
curl.exe -i -X POST http://127.0.0.1:8081/jobs
curl.exe -i -X POST http://127.0.0.1:8082/jobs
curl.exe -i -X POST http://127.0.0.1:8083/jobs
```

URL bawaan mengarah ke `https://jsonplaceholder.typicode.com/todos/1`. Atur `EXTERNAL_URL` sebelum menjalankan worker untuk memakai layanan lain. Alamat server bisa diatur lewat `ADDR`.

```powershell
$env:EXTERNAL_URL = 'http://127.0.0.1:9000/todos/1'
go run ./gin
```

Setiap panggilan keluar punya batas waktu 5 detik dan ukuran respons maksimal 1 MiB. Antrean menyimpan hingga 8 job yang menunggu. Server mengembalikan `503 Service Unavailable` jika antrean penuh. Saat proses dihentikan, job yang masih mengantre akan hilang karena antreannya disimpan di memori.
