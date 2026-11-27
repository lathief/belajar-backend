<?php

function bacaJumlah(string $teks): ?int
{
    $jumlah = filter_var($teks, FILTER_VALIDATE_INT);
    if ($jumlah === false || $jumlah < 1 || $jumlah > 5) {
        return null;
    }

    return $jumlah;
}
