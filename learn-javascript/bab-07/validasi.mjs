export function bacaJumlah(teks) {
  // Hanya izinkan digit desimal sebelum mengonversi teks ke angka.
  if (!/^\d+$/.test(teks)) return null;
  const n = Number(teks);
  return n >= 1 && n <= 5 ? n : null;
}
