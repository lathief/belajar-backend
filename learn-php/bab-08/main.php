<?php

function parseJumlah(string $teks): int
{
    $jumlah = filter_var($teks, FILTER_VALIDATE_INT);
    if ($jumlah === false) {
        throw new InvalidArgumentException("input_tidak_valid");
    }
    if ($jumlah < 1 || $jumlah > 5) {
        throw new DomainException("jumlah_di_luar_batas");
    }

    return $jumlah;
}

foreach (["2", "dua", "0"] as $teks) {
    try {
        echo parseJumlah($teks) . "\n";
    } catch (InvalidArgumentException | DomainException $error) {
        echo $error->getMessage() . "\n";
    }
}
