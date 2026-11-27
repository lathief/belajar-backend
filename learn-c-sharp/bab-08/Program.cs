#:property PublishAot=false
using System;

static int? BacaJumlah(string teks)
{
    // n baru dipakai ketika konversi berhasil.
    if (int.TryParse(teks, out int n) && n >= 1 && n <= 5) return n;
    return null;
}

foreach (string teks in new[] { "2", "0", "x" })
{
    int? jumlah = BacaJumlah(teks);
    Console.WriteLine(jumlah?.ToString() ?? "ditolak");
}
