#:property PublishAot=false
using System;

Console.WriteLine("hasil=" + Kirim([503, 503, 200], 4));
Console.WriteLine("hasil=" + Kirim([503, 503, 200], 3));
Console.WriteLine("hasil=" + Kirim([400, 200], 4));

static string Kirim(int[] respons, int deadline)
{
    int waktu = 0;
    int jeda = 1;
    for (int i = 0; i < respons.Length; i++)
    {
        int kode = respons[i];
        Console.WriteLine($"t={waktu} percobaan={i + 1} kode={kode}");
        if (kode == 200) return "sukses";
        if (kode != 503) return "permanen";
        if (i + 1 == respons.Length) return "habis";
        // Sediakan waktu untuk percobaan berikutnya sebelum menjadwalkannya.
        if (waktu + jeda >= deadline) return "deadline";
        waktu += jeda;
        jeda *= 2;
    }
    return "habis";
}
