class Alamat {
  kota = "Bandung";
  versi = 1;

  ubah(versiDibaca, kota) {
    if (this.versi !== versiDibaca) return "konflik";
    // Method sinkron menyelesaikan kedua perubahan dalam satu pemanggilan.
    this.kota = kota;
    this.versi++;
    return "tersimpan";
  }
}

const alamat = new Alamat();
const bacaA = alamat.versi;
const bacaB = alamat.versi;
function lapor(pelaku, versi, kota) {
  const hasil = alamat.ubah(versi, kota);
  console.log(`${pelaku} ${hasil} kota=${alamat.kota} versi=${alamat.versi}`);
}
lapor("A", bacaA, "Solo");
lapor("B", bacaB, "Bogor");
// B membaca ulang dan pengguna mengonfirmasi perubahan.
lapor("B", alamat.versi, "Bogor");
