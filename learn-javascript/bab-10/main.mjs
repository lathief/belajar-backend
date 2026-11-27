class Gudang {
  #stok;
  #riwayat = new Map();

  constructor(stokAwal) {
    this.#stok = stokAwal;
  }

  get stok() {
    return this.#stok;
  }

  pesan(id, jumlah) {
    if (!Number.isInteger(jumlah) || jumlah < 1 || jumlah > 5) {
      return "jumlah_invalid";
    }

    // Objek bukti selalu terisi untuk setiap ID yang sudah diproses.
    const awal = this.#riwayat.get(id);
    if (awal !== undefined) {
      if (awal.jumlah !== jumlah) return "konflik";
      return awal.hasil;
    }

    let hasil = "stok_kurang";
    if (jumlah <= this.#stok) {
      this.#stok -= jumlah;
      hasil = "diterima";
    }
    this.#riwayat.set(id, { jumlah, hasil });
    return hasil;
  }
}

const gudang = new Gudang(5);
const urutan = [
  ["A", 3], ["B", 3], ["C", 0], ["C", 2],
  ["A", 3], ["A", 1], ["B", 3],
];
for (const [id, jumlah] of urutan) {
  const hasil = gudang.pesan(id, jumlah);
  console.log(`${id} ${jumlah} ${hasil} stok=${gudang.stok}`);
}
