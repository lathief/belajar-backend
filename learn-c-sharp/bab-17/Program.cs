#:property PublishAot=false
using System;

var tugas = new Tugas();
int a = tugas.Klaim(0) ?? throw new InvalidOperationException();
Console.WriteLine("klaim_dini=" + tugas.Klaim(4).HasValue.ToString().ToLowerInvariant());
Console.WriteLine("batas=" + tugas.Simpan(a, 5));
int b = tugas.Klaim(5) ?? throw new InvalidOperationException();
Console.WriteLine("lama=" + tugas.Simpan(a, 6));
Console.WriteLine("baru=" + tugas.Simpan(b, 6));
Console.WriteLine("ulang=" + tugas.Simpan(b, 7));

class Tugas
{
    private int generasi;
    private int kedaluwarsa;
    private bool selesai;

    public int? Klaim(int kini)
    {
        if (selesai || kini < kedaluwarsa) return null;
        kedaluwarsa = kini + 5;
        return ++generasi;
    }

    public string Simpan(int token, int kini)
    {
        // Tepat pada kedaluwarsa pemilik lama sudah kehilangan hak simpan.
        if (token == 0 || selesai || token != generasi || kini >= kedaluwarsa) return "ditolak";
        selesai = true;
        return "tersimpan";
    }
}
