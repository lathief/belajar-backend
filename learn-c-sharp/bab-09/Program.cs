#:property PublishAot=false
using System;

long unit = 0;
int ditolak = 0;
foreach (string teks in args)
{
    // Pola rentang memeriksa nilai sesudah konversi berhasil.
    if (int.TryParse(teks, out int n) && n is >= 1 and <= 5)
        unit += n;
    else
        ditolak++;
}
Console.WriteLine($"unit={unit} ditolak={ditolak}");
