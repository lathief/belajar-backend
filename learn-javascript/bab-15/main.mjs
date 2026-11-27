class Ember {
  constructor(kapasitas) {
    this.kapasitas = kapasitas;
    this.token = kapasitas;
    this.terakhir = 0;
  }

  ambil(detik) {
    // Isi mengikuti waktu berlalu, termasuk ketika tak ada permintaan.
    this.token = Math.min(this.kapasitas, this.token + detik - this.terakhir);
    this.terakhir = detik;
    if (this.token === 0) return "tunda";
    this.token--;
    return "lolos";
  }
}

const ember = new Ember(2);
for (const t of [0, 0, 0, 1, 6, 6, 6]) {
  const hasil = ember.ambil(t);
  console.log(`t=${t} ${hasil} token=${ember.token}`);
}
