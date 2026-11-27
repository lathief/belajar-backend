import { readFile } from "node:fs/promises";

async function bacaPesanan() {
  // URL relatif ke modul agar fixture ditemukan dari direktori kerja mana pun.
  const lokasi = new URL("./pesanan.json", import.meta.url);
  const teks = await readFile(lokasi, "utf8");
  return JSON.parse(teks);
}

try {
  const pesanan = await bacaPesanan();
  console.log(`${pesanan.kode} ${pesanan.status}`);
} catch (error) {
  console.error("Gagal membaca pesanan:", error.message);
  process.exitCode = 1;
}
