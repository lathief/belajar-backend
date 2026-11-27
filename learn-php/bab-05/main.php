<?php

$stok = [
    "PS-1" => 0,
    "KB-2" => 4,
];

foreach (["PS-1", "KODE-HILANG"] as $kode) {
    $nilai = $stok[$kode] ?? "tidak_tercatat";
    echo "{$kode} stok={$nilai}\n";
}
