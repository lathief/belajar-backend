#:property PublishAot=false
using System;

static decimal Subtotal(decimal harga, int jumlah)
{
    // Pemanggil memilih bagaimana hasil akan ditampilkan.
    return harga * jumlah;
}

Console.WriteLine($"total={Subtotal(5000m, 3)}");
