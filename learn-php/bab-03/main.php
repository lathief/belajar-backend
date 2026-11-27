<?php

$jumlahPesanan = [0, 1, 2, 6, 5];
$diterima = 0;

foreach ($jumlahPesanan as $jumlah) {
    if ($jumlah >= 1 && $jumlah <= 5) {
        $diterima++;
    }
}

echo "jumlah_diterima={$diterima}\n";
