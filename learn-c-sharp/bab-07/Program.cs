#:property PublishAot=false
using System;
using System.Collections.Generic;
using System.Linq;

var jumlah = new List<int> { 2, 0, 3 };
// Sum mengevaluasi rangkaian elemen yang lolos filter.
var laporan = new Dictionary<string, int>();
laporan["unit"] = jumlah.Where(n => n > 0).Sum();
Console.WriteLine("unit=" + laporan["unit"]);
