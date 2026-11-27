function simulasi(responsHilang, relayPutus, ackHilang) {
  let produsen = { stok: 5, dipesan: false, tertunda: false };
  function pesan() {
    if (!produsen.dipesan) {
      // Pergantian state memodelkan satu commit pesanan dan outbox.
      produsen = { stok: produsen.stok - 1, dipesan: true, tertunda: true };
    }
  }
  pesan();
  if (responsHilang) pesan();

  const antrean = [];
  let publikasi = 0;
  while (produsen.tertunda) {
    antrean.push("A");
    publikasi++;
    if (relayPutus) {
      relayPutus = false;
      continue;
    }
    produsen.tertunda = false;
  }

  const inbox = new Set();
  let paket = 0;
  while (antrean.length > 0) {
    const id = antrean[0];
    // Dua perubahan ini satu transaksi konsumen dalam model.
    if (!inbox.has(id)) {
      inbox.add(id);
      paket++;
    }
    if (ackHilang) {
      ackHilang = false;
      continue;
    }
    antrean.shift();
  }
  return { stok: produsen.stok, publikasi, paket };
}

let lulus = 0;
for (const respons of [false, true]) {
  for (const relay of [false, true]) {
    for (const ack of [false, true]) {
      const { stok, publikasi, paket } = simulasi(respons, relay, ack);
      if (stok !== 4 || publikasi !== (relay ? 2 : 1) || paket !== 1) {
        throw new Error("Invarian gagal");
      }
      lulus++;
      if (respons && relay && ack) console.log(`stok=${stok} publikasi=${publikasi} paket=${paket}`);
    }
  }
}
console.log(`skenario_lulus=${lulus}`);
