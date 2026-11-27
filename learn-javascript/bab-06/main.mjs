const jumlah = [2, 0, 3];
const nilai = jumlah.filter(n => n > 0).map(n => n * 5000);
// Nilai awal memastikan array kosong menghasilkan total nol.
const total = nilai.reduce((akumulasi, harga) => akumulasi + harga, 0);
const laporan = new Map([["omzet", total]]);
console.log(`omzet=${laporan.get("omzet")}`);
