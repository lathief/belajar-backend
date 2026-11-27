<?php

$deadline = 5;
$waktu = 0;
$jeda = 2;
$percobaan = 0;
$berhasil = false;

while ($waktu < $deadline) {
    $percobaan++;
    echo "percobaan={$percobaan} gagal\n";
    if ($waktu + $jeda >= $deadline) {
        break;
    }
    $waktu += $jeda;
    $jeda *= 2;
}

if (!$berhasil) {
    echo "deadline_tercapai\n";
}
