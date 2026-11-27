const daftar = [0, 1, 4, 8];
let valid = 0;
for (const n of daftar) {
  // Rentang ini mencakup kedua batas.
  if (n >= 1 && n <= 5) {
    valid++;
  }
}
console.log(`valid=${valid}`);
