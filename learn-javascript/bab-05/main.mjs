const awal = { kode: "PS-1", status: "baru" };
const { kode } = awal;
// Salin field terlebih dahulu, lalu ganti status pada salinan.
const akhir = { ...awal, status: "lunas" };
console.log(`${kode} ${awal.status} -> ${akhir.status}`);
