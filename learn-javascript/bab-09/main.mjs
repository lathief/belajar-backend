import { bacaJumlah } from "../bab-07/validasi.mjs";

let unit = 0;
let ditolak = 0;
for (const teks of process.argv.slice(2)) {
  const jumlah = bacaJumlah(teks);
  // Validator mengembalikan null untuk pesanan yang ditolak.
  if (jumlah === null) ditolak++;
  else unit += jumlah;
}
console.log(`unit=${unit} ditolak=${ditolak}`);
