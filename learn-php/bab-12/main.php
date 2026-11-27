<?php

final class StokBerversi
{
    public function __construct(private int $jumlah, private int $versi = 1)
    {
    }

    public function simpan(int $versiDiharapkan, int $jumlahBaru): bool
    {
        if ($versiDiharapkan !== $this->versi) {
            return false;
        }
        $this->jumlah = $jumlahBaru;
        $this->versi++;
        return true;
    }

    public function jumlah(): int
    {
        return $this->jumlah;
    }

    public function versi(): int
    {
        return $this->versi;
    }
}

$stok = new StokBerversi(10);
$versiDibacaKlienA = $stok->versi();
$versiDibacaKlienB = $stok->versi();
if ($stok->simpan($versiDibacaKlienA, 8)) {
    echo "tersimpan versi=" . $stok->versi() . "\n";
}
if (!$stok->simpan($versiDibacaKlienB, 6)) {
    echo "konflik versi_terkini=" . $stok->versi() . "\n";
}
echo "stok=" . $stok->jumlah() . " versi=" . $stok->versi() . "\n";
