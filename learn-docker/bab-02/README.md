# Bab 02 — Port dan siklus hidup server

Satu container bisa terus melayani permintaan selama proses utamanya berjalan. Di sini BusyBox `httpd` berjalan di foreground dengan `-f`, sedangkan Docker menjalankan container di latar belakang dengan `-d`.

## Nyalakan server

```powershell
docker run -d --name ld02-web -p 127.0.0.1:8072:8080 busybox:1.37 sh -c 'mkdir -p /srv && echo Katalog > /srv/index.html && echo Server-siap && exec httpd -f -p 8080 -h /srv'
docker port ld02-web
curl.exe -s http://127.0.0.1:8072/
docker logs ld02-web
```

Browser atau curl menghubungi port `8072` pada komputer. Docker meneruskannya ke port `8080` di dalam container. Respons HTTP berisi `Katalog`, sementara log menampilkan `Server-siap` dari perintah startup.

## Latihan

```powershell
docker stop ld02-web
docker ps -a --filter name=ld02-web
docker start ld02-web
docker exec ld02-web cat /srv/index.html
```

Periksa bahwa status menjadi `Exited` setelah server dihentikan dan halaman bisa diakses lagi sesudah `start`. Perintah `exec` menjalankan proses tambahan di dalam container yang aktif.

## Bereskan

```powershell
docker stop ld02-web
docker rm ld02-web
```

[Bab sebelumnya](../bab-01/README.md) · [Roadmap](../README.md) · [Bab berikutnya](../bab-03/README.md)
