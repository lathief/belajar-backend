#:property PublishAot=false
using System;

var status = Status.Baru;
foreach (string aksi in new[] { "kirim", "bayar", "batal", "kirim", "bayar" })
{
    var (berikut, ok) = Transisi(status, aksi);
    Console.WriteLine($"{status.ToString().ToLowerInvariant()} {aksi} -> {berikut.ToString().ToLowerInvariant()} ok={ok.ToString().ToLowerInvariant()}");
    status = berikut;
}

static (Status, bool) Transisi(Status awal, string aksi) => (awal, aksi) switch
{
    (Status.Baru, "bayar") => (Status.Lunas, true),
    (Status.Baru, "batal") => (Status.Batal, true),
    (Status.Lunas, "kirim") => (Status.Dikirim, true),
    _ => (awal, false) // Penolakan mempertahankan keadaan semula.
};

enum Status { Baru, Lunas, Dikirim, Batal }
