#:property PublishAot=false
using System;

var alamat = new Alamat();
int bacaA = alamat.Versi;
int bacaB = alamat.Versi;
Lapor("A", bacaA, "Solo");
Lapor("B", bacaB, "Bogor");
Lapor("B", alamat.Versi, "Bogor"); // Pengguna mengonfirmasi sesudah membaca ulang.

void Lapor(string pelaku, int versi, string kota)
{
    string hasil = alamat.Ubah(versi, kota);
    Console.WriteLine($"{pelaku} {hasil} kota={alamat.Kota} versi={alamat.Versi}");
}

class Alamat
{
    public string Kota { get; private set; } = "Bandung";
    public int Versi { get; private set; } = 1;

    public string Ubah(int dibaca, string kota)
    {
        if (Versi != dibaca) return "konflik";
        Kota = kota;
        Versi++;
        return "tersimpan";
    }
}
