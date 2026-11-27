#:property PublishAot=false
using System;

var stok = new Persediaan(2);
stok.AmbilSatu();
Console.WriteLine($"sisa={stok.Sisa}");

class Persediaan
{
    public int Sisa { get; private set; }

    public Persediaan(int awal) => Sisa = awal;

    public void AmbilSatu()
    {
        // Property hanya diperbarui ketika barang tersedia.
        if (Sisa > 0) Sisa--;
    }
}
