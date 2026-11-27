#:property PublishAot=false
using System;

int[] daftar = { 0, 1, 4, 8 };
int valid = 0;
foreach (int n in daftar)
{
    // Kedua batas termasuk dalam rentang pesanan sah.
    if (n >= 1 && n <= 5)
    {
        valid++;
    }
}
Console.WriteLine($"valid={valid}");
