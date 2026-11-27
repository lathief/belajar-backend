class Ringkasan {
  posisi = 0;
  stok = 0;

  terapkan({ nomor, delta }) {
    if (nomor <= this.posisi) return "duplikat";
    if (nomor !== this.posisi + 1) return "celah";
    // Nomor berikutnya baru diakui sesudah delta masuk ke ringkasan.
    this.stok += delta;
    this.posisi = nomor;
    return "diterapkan";
  }
}

const r = new Ringkasan();
const log = [
  { nomor: 1, delta: 5 }, { nomor: 3, delta: -2 },
  { nomor: 2, delta: -1 }, { nomor: 2, delta: -1 }, { nomor: 3, delta: -2 },
];
for (const e of log) {
  const hasil = r.terapkan(e);
  console.log(`${e.nomor} ${hasil} posisi=${r.posisi} stok=${r.stok}`);
}
