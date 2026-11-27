#:property PublishAot=false
using System;
using System.Collections.Generic;

bool[] opsi = [false, true];
int lulus = 0;
foreach (bool respons in opsi)
{
    foreach (bool relay in opsi)
    {
        foreach (bool ack in opsi)
        {
            var hasil = Simulasi(respons, relay, ack);
            if (hasil != (4, relay ? 2 : 1, 1)) throw new InvalidOperationException("Invarian gagal");
            lulus++;
            if (respons && relay && ack)
                Console.WriteLine($"stok={hasil.Stok} publikasi={hasil.Publikasi} paket={hasil.Paket}");
        }
    }
}
Console.WriteLine($"skenario_lulus={lulus}");

static (int Stok, int Publikasi, int Paket) Simulasi(bool responsHilang, bool relayPutus, bool ackHilang)
{
    var p = new Produsen(5, false, false).Pesan();
    if (responsHilang) p = p.Pesan();

    var antrean = new Queue<string>();
    int publikasi = 0;
    while (p.Tertunda)
    {
        antrean.Enqueue("A");
        publikasi++;
        if (relayPutus)
        {
            relayPutus = false;
            continue;
        }
        p = p with { Tertunda = false };
    }

    var inbox = new HashSet<string>();
    int paket = 0;
    while (antrean.Count > 0)
    {
        // Add melaporkan apakah ID baru, sebelum pesan diakui melalui Dequeue.
        if (inbox.Add(antrean.Peek())) paket++;
        if (ackHilang)
        {
            ackHilang = false;
            continue;
        }
        antrean.Dequeue();
    }
    return (p.Stok, publikasi, paket);
}

record Produsen(int Stok, bool Dipesan, bool Tertunda)
{
    public Produsen Pesan() => Dipesan ? this : this with { Stok = Stok - 1, Dipesan = true, Tertunda = true };
}
