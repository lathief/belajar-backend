# Bab 04 — Mengedit halaman dengan bind mount

Saat mengembangkan halaman, kita bisa memasang folder proyek ke container agar perubahan langsung terlihat. Sumber bind mount di bab ini adalah [folder katalog](katalog/index.html) di komputer.

## Pasang folder

```powershell
$folder = (Resolve-Path ./katalog).Path
docker run -d --name ld04-web -p 127.0.0.1:8074:8080 --mount "type=bind,source=$folder,target=/srv,readonly" busybox:1.37 httpd -f -p 8080 -h /srv
curl.exe -s http://127.0.0.1:8074/
```

Cari teks `Stok buku: 3`. Nama `/srv` adalah lokasi folder ketika dibaca dari dalam container. Opsi `readonly` memberi server akses baca ke folder tersebut.

## Latihan

Buka `katalog/index.html` di editor, ubah angka stok menjadi 4, lalu simpan. Ulangi perintah curl. Respons langsung memuat angka baru dan `docker exec ld04-web cat /srv/index.html` memperlihatkan isi berkas yang sama.

## Bereskan

```powershell
docker stop ld04-web
docker rm ld04-web
```

Kembalikan angka stok ke 3 jika ingin mengulang latihan dari awal.

[Bab sebelumnya](../bab-03/README.md) · [Roadmap](../README.md) · [Bab berikutnya](../bab-05/README.md)
