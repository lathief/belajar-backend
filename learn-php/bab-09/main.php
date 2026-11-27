<?php

require_once __DIR__ . "/../bab-07/validasi.php";

$unit = 0;
$ditolak = 0;
foreach (array_slice($argv, 1) as $teks) {
    $jumlah = bacaJumlah($teks);
    if ($jumlah === null) {
        $ditolak++;
        continue;
    }
    $unit += $jumlah;
}

echo "unit={$unit} ditolak={$ditolak}\n";
