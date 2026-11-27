#:property PublishAot=false
using System;

const decimal harga = 5000m;
int jumlah = 2;
// Jumlah pesanan masih dapat berubah sebelum dihitung.
jumlah += 1;
Console.WriteLine($"subtotal={harga * jumlah}");
