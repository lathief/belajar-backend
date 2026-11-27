function subtotal(harga, jumlah) {
  // Simpan rumus di satu tempat agar dapat diuji terpisah.
  return harga * jumlah;
}

console.log(`total=${subtotal(5000, 3)}`);
