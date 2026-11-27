#:property PublishAot=false
using System;

var ember = new Ember(2);
foreach (int t in new[] { 0, 0, 0, 1, 6, 6, 6 })
{
    string hasil = ember.Ambil(t);
    Console.WriteLine($"t={t} {hasil} token={ember.Token}");
}

class Ember(int kapasitas)
{
    public int Token { get; private set; } = kapasitas;
    private readonly int batasToken = kapasitas;
    private int terakhir;

    public string Ambil(int detik)
    {
        // Token tersisa ditambah selisih waktu hingga batas kapasitas.
        Token = Math.Min(batasToken, Token + detik - terakhir);
        terakhir = detik;
        if (Token == 0) return "tunda";
        Token--;
        return "lolos";
    }
}
