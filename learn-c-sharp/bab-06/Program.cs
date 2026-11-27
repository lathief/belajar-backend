#:property PublishAot=false
using System;

var awal = new Pesanan("PS-1", Status.Baru);
var akhir = awal with { Status = Status.Bayar };
// Pola cadangan menangani nilai enum numerik lain.
string label = akhir.Status switch
{
    Status.Baru => "menunggu",
    Status.Bayar => "lunas",
    _ => "tak dikenal"
};
Console.WriteLine($"{akhir.Kode} {label}");

enum Status { Baru, Bayar }
record Pesanan(string Kode, Status Status);
