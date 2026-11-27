#:property PublishAot=false
using System;
using System.Collections.Generic;
using System.Linq;

var log = new List<int> { 101, 104, 107, 112 };
int batas = log[^1];
int cursor = 0;
while (true)
{
    var ids = Halaman(log, cursor, batas, 2);
    if (ids.Length == 0) break;
    Console.WriteLine("halaman=" + string.Join(",", ids));
    if (cursor == 0) log.Add(115); // Batas sesi tetap 112.
    cursor = ids[^1];
}
Console.WriteLine("susulan=" + string.Join(",", Halaman(log, cursor, log[^1], 2)));

static int[] Halaman(List<int> log, int setelah, int batas, int ukuran)
{
    return log.SkipWhile(id => id <= setelah)
        .TakeWhile(id => id <= batas)
        .Take(ukuran)
        .ToArray();
}
