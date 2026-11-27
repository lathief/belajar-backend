class Tugas {
  generasi = 0;
  kedaluwarsa = 0;
  selesai = false;

  klaim(kini) {
    if (this.selesai || kini < this.kedaluwarsa) return null;
    this.kedaluwarsa = kini + 5;
    return ++this.generasi;
  }

  simpan(token, kini) {
    // Token yang masih sama pun gugur tepat pada akhir lease.
    if (!token || this.selesai || token !== this.generasi || kini >= this.kedaluwarsa) {
      return "ditolak";
    }
    this.selesai = true;
    return "tersimpan";
  }
}

const tugas = new Tugas();
const a = tugas.klaim(0);
console.log("klaim_dini=" + (tugas.klaim(4) !== null));
console.log("batas=" + tugas.simpan(a, 5));
const b = tugas.klaim(5);
console.log("lama=" + tugas.simpan(a, 6));
console.log("baru=" + tugas.simpan(b, 6));
console.log("ulang=" + tugas.simpan(b, 7));
