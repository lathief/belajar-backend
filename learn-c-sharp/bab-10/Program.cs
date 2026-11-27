#:property PublishAot=false
using System;
using System.Collections.Generic;

var gudang = new Gudang(5);
(string Id, int Jumlah)[] urutan = {
    ("A", 3), ("B", 3), ("C", 0), ("C", 2),
    ("A", 3), ("A", 1), ("B", 3)
};
foreach (var (id, jumlah) in urutan)
{
    string hasil = gudang.Pesan(id, jumlah);
    Console.WriteLine($"{id} {jumlah} {hasil} stok={gudang.Stok}");
}

record Bukti(int Jumlah, string Hasil);

class Gudang(int stokAwal)
{
    public int Stok { get; private set; } = stokAwal;
    private readonly Dictionary<string, Bukti> riwayat = new();

    public string Pesan(string id, int jumlah)
    {
        if (jumlah is < 1 or > 5) return "jumlah_invalid";

        // TryGetValue mengambil bukti dan memeriksa keberadaannya sekaligus.
        if (riwayat.TryGetValue(id, out var awal))
        {
            if (awal.Jumlah != jumlah) return "konflik";
            return awal.Hasil;
        }

        string hasil = "stok_kurang";
        if (jumlah <= Stok)
        {
            Stok -= jumlah;
            hasil = "diterima";
        }
        riwayat.Add(id, new Bukti(jumlah, hasil));
        return hasil;
    }
}
