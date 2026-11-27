function halaman(log, setelah, batas, ukuran) {
  const hasil = [];
  for (const id of log) {
    if (id > batas || hasil.length === ukuran) break;
    if (id > setelah) hasil.push(id);
  }
  return hasil;
}

const log = [101, 104, 107, 112];
const batas = log.at(-1);
let cursor = 0;
while (true) {
  const ids = halaman(log, cursor, batas, 2);
  if (ids.length === 0) break;
  console.log(`halaman=${ids.join(",")}`);
  if (cursor === 0) log.push(115); // Sesi ini sudah mengunci batas atasnya.
  cursor = ids.at(-1);
}
const susulan = halaman(log, cursor, log.at(-1), 2);
console.log(`susulan=${susulan.join(",")}`);
