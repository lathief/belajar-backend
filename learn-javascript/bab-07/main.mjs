import { bacaJumlah } from "./validasi.mjs";

for (const teks of ["2", "x"]) {
  // Nullish coalescing memberi label saat validator mengembalikan null.
  console.log(bacaJumlah(teks) ?? "ditolak");
}
