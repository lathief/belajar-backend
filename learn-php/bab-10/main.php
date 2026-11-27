<?php

final class Gudang
{
    private array $riwayat = [];

    public function __construct(private int $stok)
    {
    }

    public function stok(): int
    {
        return $this->stok;
    }

    public function pesan(string $id, int $jumlah): string
    {
        if ($jumlah < 1 || $jumlah > 5) {
            return "jumlah_invalid";
        }
        if (isset($this->riwayat[$id])) {
            $awal = $this->riwayat[$id];
            return $awal["jumlah"] === $jumlah ? $awal["hasil"] : "konflik";
        }

        $hasil = "stok_kurang";
        if ($jumlah <= $this->stok) {
            $this->stok -= $jumlah;
            $hasil = "diterima";
        }
        $this->riwayat[$id] = ["jumlah" => $jumlah, "hasil" => $hasil];
        return $hasil;
    }
}

$gudang = new Gudang(5);
$urutan = [
    ["A", 3], ["B", 3], ["C", 0], ["C", 2],
    ["A", 3], ["A", 1], ["B", 3],
];
foreach ($urutan as [$id, $jumlah]) {
    $hasil = $gudang->pesan($id, $jumlah);
    echo "{$id} {$jumlah} {$hasil} stok=" . $gudang->stok() . "\n";
}
