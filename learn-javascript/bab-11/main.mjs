function transisi(awal, aksi) {
  // Daftar ini mendefinisikan seluruh jalur yang diizinkan.
  const berikut = aturan.get(awal)?.get(aksi);
  return berikut === undefined
    ? { status: awal, ok: false }
    : { status: berikut, ok: true };
}

const aturan = new Map([
  ["baru", new Map([["bayar", "lunas"], ["batal", "batal"]])],
  ["lunas", new Map([["kirim", "dikirim"]])],
]);

let status = "baru";
for (const aksi of ["kirim", "bayar", "batal", "kirim", "bayar"]) {
  const hasil = transisi(status, aksi);
  console.log(`${status} ${aksi} -> ${hasil.status} ok=${hasil.ok}`);
  status = hasil.status;
}
