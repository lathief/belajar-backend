<?php

function hitungTotal(int $harga, int $jumlah): int
{
    return $harga * $jumlah;
}

$total = hitungTotal(4000, 3);
echo "total={$total}\n";
