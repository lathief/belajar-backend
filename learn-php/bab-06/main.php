<?php

final class Gudang
{
    public function __construct(private int $stok)
    {
    }

    public function stok(): int
    {
        return $this->stok;
    }

    public function pesan(int $jumlah): string
    {
        if ($jumlah < 1) {
            return "jumlah_invalid";
        }
        if ($jumlah > $this->stok) {
            return "stok_kurang";
        }

        $this->stok -= $jumlah;
        return "diterima";
    }
}

$gudang = new Gudang(5);
foreach ([2, 4] as $jumlah) {
    $hasil = $gudang->pesan($jumlah);
    echo $hasil . " stok=" . $gudang->stok() . "\n";
}
