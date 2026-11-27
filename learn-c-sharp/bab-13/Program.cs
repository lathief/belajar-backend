#:property PublishAot=false
using System;

var r = new Ringkasan();
(int Nomor, int Delta)[] log = { (1, 5), (3, -2), (2, -1), (2, -1), (3, -2) };
foreach (var (nomor, delta) in log)
{
    string hasil = r.Terapkan(nomor, delta);
    Console.WriteLine($"{nomor} {hasil} posisi={r.Posisi} stok={r.Stok}");
}

class Ringkasan
{
    public int Posisi { get; private set; }
    public int Stok { get; private set; }

    public string Terapkan(int nomor, int delta)
    {
        if (nomor <= Posisi) return "duplikat";
        if (nomor != Posisi + 1) return "celah";
        // Simpan posisi untuk delta terakhir yang sudah diterapkan.
        Stok += delta;
        Posisi = nomor;
        return "diterapkan";
    }
}
