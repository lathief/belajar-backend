function kirim(respons, deadline) {
  let waktu = 0;
  let jeda = 1;
  for (const [i, kode] of respons.entries()) {
    console.log(`t=${waktu} percobaan=${i + 1} kode=${kode}`);
    if (kode === 200) return "sukses";
    if (kode !== 503) return "permanen";
    if (i + 1 === respons.length) return "habis";
    // Periksa jadwal berikutnya sebelum memakai kesempatan retry.
    if (waktu + jeda >= deadline) return "deadline";
    waktu += jeda;
    jeda *= 2;
  }
}

for (const [respons, deadline] of [
  [[503, 503, 200], 4], [[503, 503, 200], 3], [[400, 200], 4],
]) {
  console.log("hasil=" + kirim(respons, deadline));
}
